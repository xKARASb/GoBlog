package routers

import (
	"net/http"

	"github.com/xkarasb/blog/internal/transport/http/handlers"
)

func GetReaderRouter(parentMux *http.ServeMux) *http.ServeMux {
	controller := handlers.NewReaderController()
	router := http.NewServeMux()

	parentMux.HandleFunc("GET /posts", controller.ViewSelectionHandler)

	return router
}
