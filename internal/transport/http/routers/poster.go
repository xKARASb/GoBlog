package routers

import (
	"net/http"

	"github.com/xkarasb/blog/internal/transport/http/handlers"
)

func GetPosterRouter(parentMux *http.ServeMux) *http.ServeMux {
	controller := handlers.NewPosterController()
	router := http.NewServeMux()

	parentMux.HandleFunc("POST /posts", controller.CreatePostHandler)
	router.HandleFunc("POST /posts/{postId}/images", controller.AddImageHandler)
	router.HandleFunc("PUT /posts/{postId}", controller.AddImageHandler)
	router.HandleFunc("DELTE /posts/{postId}/images/{imageId}", controller.DeleteImageHandler)
	router.HandleFunc("DELTE /posts/{postId}/status", controller.PublishHandler)

	return router
}
