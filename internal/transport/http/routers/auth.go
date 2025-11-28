package routers

import (
	"net/http"

	"github.com/xkarasb/blog/internal/transport/http/handlers"
)

func GetAuthRouter(parentMux *http.ServeMux) *http.ServeMux {
	controller := handlers.NewAuthController()
	router := http.NewServeMux()

	router.HandleFunc("POST /auth/register", controller.RefreshHandler)
	router.HandleFunc("POST /auth/login", controller.LoginHandler)
	router.HandleFunc("POST /auth/refresh-token", controller.RefreshHandler)

	return router
}
