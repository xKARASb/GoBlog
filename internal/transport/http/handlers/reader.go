package handlers

import (
	"fmt"
	"net/http"

	"github.com/xkarasb/blog/internal/core/dto"
)

type ReaderController struct{}

func NewReaderController() *ReaderController {
	return &ReaderController{}
}

func (c *ReaderController) ViewSelectionHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userRaw := ctx.Value("user")
	user, ok := userRaw.(*dto.UserDB)
	if !ok {
		return
	}
	switch user.Role {
	case "author":
		c.authorView(w, r)
	case "reader":
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
	fmt.Fprintf(w, "Its create post!\n")
}
