package handlers

import (
	"fmt"
	"net/http"
)

type ReaderController struct{}

func NewReaderController() *ReaderController {
	return &ReaderController{}
}

func (c *ReaderController) ViewSelectionHandler(w http.ResponseWriter, r *http.Request) {
	c.readerView(w, r)
}

func (c *ReaderController) readerView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!\n")
}

func (c *ReaderController) authorView(w http.ResponseWriter, r *http.Request) {
}
