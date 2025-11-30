package routers

import (
	"net/http"

	"github.com/xkarasb/blog/internal/core/repository"
	"github.com/xkarasb/blog/internal/core/service"
	"github.com/xkarasb/blog/internal/transport/http/handlers"
	"github.com/xkarasb/blog/pkg/db/postgres"
)

func GetAuthRouter(db *postgres.DB) (*http.ServeMux, *service.AuthService) {
	repo := repository.NewBlogRepository(db)
	serv := service.NewAuthService(repo, "secret")
	controller := handlers.NewAuthController(serv)
	router := http.NewServeMux()

	router.HandleFunc("POST /auth/register", controller.RegisterHandler)
	router.HandleFunc("POST /auth/login", controller.LoginHandler)
	router.HandleFunc("POST /auth/refresh-token", controller.RefreshHandler)

	return router, serv
}
