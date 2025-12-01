package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/xkarasb/blog/internal/core/dto"
	"github.com/xkarasb/blog/pkg/errors"
	"github.com/xkarasb/blog/pkg/types"
)

type ReaderService interface {
	NewPost(authorId uuid.UUID, post *dto.CreatePostRequest) (*dto.CreatePostResponse, error)
}

type ReaderController struct {
	service ReaderService
}

func NewReaderController(service ReaderService) *ReaderController {
	return &ReaderController{
		service: service,
	}
}

func (c *ReaderController) ViewSelectionHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, ok := ctx.Value(types.CtxUser).(*dto.UserDB)
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Incorrect user")
		return
	}
	switch user.Role {
	case types.Author:
		c.authorView(w, r)
	case types.Reader:
		c.readerView(w, r)
	default:
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Incorrect user")
	}
}

func (c *ReaderController) readerView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Its reader!\n")

}

func (c *ReaderController) authorView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Its author!\n")
}

func (c *ReaderController) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, ok := ctx.Value(types.CtxUser).(*dto.UserDB)
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Incorrect user")
		return
	}

	reqPost := &dto.CreatePostRequest{}
	if err := json.NewDecoder(r.Body).Decode(reqPost); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Incorrect body")
		return
	}
	resPost, err := c.service.NewPost(user.UserId, reqPost)

	if err != nil {
		if err == errors.ErrorKeyIdempotencyAlreadyUsed {
			w.WriteHeader(http.StatusConflict)
			fmt.Fprintln(w, "Idempotency key already used")
			return
		}
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintln(w, "Something wrong")
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resPost)

}
