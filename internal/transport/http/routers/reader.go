package routers

import (
	"net/http"

	"github.com/xkarasb/blog/internal/core/service"
	"github.com/xkarasb/blog/internal/transport/http/handlers"
	"github.com/xkarasb/blog/internal/transport/http/middlewares"
)

func GetReaderRouter(service *service.ReaderService, authMiddlewareManager *middlewares.AuthMiddlewareManager) *http.ServeMux {
	controller := handlers.NewReaderController(service)
	router := http.NewServeMux()

	router.HandleFunc("GET /posts", controller.ViewSelectionHandler)
	router.Handle("POST /posts", authMiddlewareManager.AuthorOnlyMiddleware(http.HandlerFunc(controller.CreatePostHandler)))

	return router
}
