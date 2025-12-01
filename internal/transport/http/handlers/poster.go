package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/xkarasb/blog/internal/core/dto"
	"github.com/xkarasb/blog/pkg/errors"
	"github.com/xkarasb/blog/pkg/types"
)

type PosterService interface {
	EditPost(userId, postId uuid.UUID, post *dto.EditPostRequest) (*dto.EditPostResponse, error)
	PublishPost(userId, postId uuid.UUID, post *dto.PublishPostRequest) (*dto.PublishPostResponse, error)
}

type PosterController struct {
	service PosterService
}

func NewPosterController(service PosterService) *PosterController {
	return &PosterController{service}
}

func (c *PosterController) AddImageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Poster %s %s\n", r.Method, r.URL)
}

func (c *PosterController) EditPostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, ok := ctx.Value(types.CtxUser).(*dto.UserDB)
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Incorrect user")
		return
	}
	reqPost := &dto.EditPostRequest{}
	if err := json.NewDecoder(r.Body).Decode(reqPost); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Incorrect body")
		return
	}

	postId, err := uuid.Parse(r.PathValue("postId"))

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Post not found")
		return
	}

	resPost, err := c.service.EditPost(user.UserId, postId, reqPost)
	if err != nil {
		switch err {
		case errors.ErrorServiceNoAccess:
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintln(w, "Access denied")
		case errors.ErrorServiceIncorrectData:
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Incorrect status")
		case sql.ErrNoRows:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "Post not found")
		default:
			w.WriteHeader(http.StatusBadGateway)
			fmt.Fprintln(w, "Something wrong")
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resPost)
}

func (c *PosterController) DeleteImageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Poster %s %s\n", r.Method, r.URL)
}

func (c *PosterController) PublishHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, ok := ctx.Value(types.CtxUser).(*dto.UserDB)
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Incorrect user")
		return
	}
	reqPost := &dto.PublishPostRequest{}
	if err := json.NewDecoder(r.Body).Decode(reqPost); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Incorrect body")
		return
	}

	postId, err := uuid.Parse(r.PathValue("postId"))

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Post not exsist")
		return
	}

	resPost, err := c.service.PublishPost(user.UserId, postId, reqPost)

	if err != nil {
		switch err {
		case errors.ErrorServiceNoAccess:
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintln(w, "Access denied")
		case errors.ErrorServiceIncorrectData:
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Incorrect status")
		case sql.ErrNoRows:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "Post not found")
		default:
			w.WriteHeader(http.StatusBadGateway)
			fmt.Fprintln(w, "Something wrong")
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resPost)
}
