package routers

import (
	"net/http"

	"github.com/xkarasb/blog/internal/core/service"
	"github.com/xkarasb/blog/internal/transport/http/handlers"
)

func GetAuthRouter(service *service.AuthService) *http.ServeMux {
	controller := handlers.NewAuthController(service)
	router := http.NewServeMux()

	router.HandleFunc("POST /auth/register", controller.RegisterHandler)
	router.HandleFunc("POST /auth/login", controller.LoginHandler)
	router.HandleFunc("POST /auth/refresh-token", controller.RefreshHandler)

	return router
}
