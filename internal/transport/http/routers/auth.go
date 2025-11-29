package routers

import (
	"net/http"

	"github.com/xkarasb/blog/internal/core/repository"
	"github.com/xkarasb/blog/internal/core/service"
	"github.com/xkarasb/blog/internal/transport/http/handlers"
	"github.com/xkarasb/blog/pkg/db/postgres"
)

func GetAuthRouter(parentMux *http.ServeMux, db *postgres.DB) *http.ServeMux {
	authRepo := repository.NewBlogRepository(db)
	authService := service.NewAuthService(authRepo)
	controller := handlers.NewAuthController(authService)
	router := http.NewServeMux()

	router.HandleFunc("POST /auth/register", controller.RegisterHandler)
	router.HandleFunc("POST /auth/login", controller.LoginHandler)
	router.HandleFunc("POST /auth/refresh-token", controller.RefreshHandler)

	return router
}
