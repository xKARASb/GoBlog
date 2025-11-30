package routers

import (
	"net/http"

	"github.com/xkarasb/blog/internal/transport/http/handlers"
)

func GetReaderRouter() *http.ServeMux {
	controller := handlers.NewReaderController()
	router := http.NewServeMux()

	router.HandleFunc("GET /posts", controller.ViewSelectionHandler)
	router.HandleFunc("POST /posts", controller.CreatePostHandler)

	return router
}
