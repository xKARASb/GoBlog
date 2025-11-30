package handlers

import (
	"fmt"
	"net/http"
)

type PosterController struct{}

func NewPosterController() *PosterController {
	return &PosterController{}
}

func (c *PosterController) AddImageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Poster %s %s\n", r.Method, r.URL)
}

func (c *PosterController) EditPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Poster %s %s\n", r.Method, r.URL)
}

func (c *PosterController) DeleteImageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Poster %s %s\n", r.Method, r.URL)
}

func (c *PosterController) PublishHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Poster %s %s\n", r.Method, r.URL)
}
