package servers

import (
	"fmt"
	"net/http"

	"github.com/xkarasb/blog/internal/core/repository"
	"github.com/xkarasb/blog/internal/core/service"
	"github.com/xkarasb/blog/internal/transport/http/middlewares"
	"github.com/xkarasb/blog/internal/transport/http/routers"
	"github.com/xkarasb/blog/pkg/db/postgres"
)

type HttpServerConfig struct {
	Address string `env:"ADDRESS" env-default:"127.0.0.1"`
	Port    int    `env:"PORT" env-default:"8080"`
}

type HttpServer struct {
	cfg  *HttpServerConfig
	http *http.Server
}

func NewHttpServer(cfg *HttpServerConfig, db *postgres.DB) *HttpServer {
	apiRouter := http.NewServeMux()

	dbRepo := repository.NewBlogRepository(db)

	authService := service.NewAuthService(dbRepo, "secret")
	readerService := service.NewReaderService(dbRepo)

	authMiddlewareManager := middlewares.NewAuthMiddlewareManager(authService)

	authRouter := routers.GetAuthRouter(authService)
	readRouter := routers.GetReaderRouter(readerService, authMiddlewareManager)
	posterRouter := routers.GetPosterRouter()

	apiRouter.Handle("/", authMiddlewareManager.AuthMiddleware(readRouter))
	// Поменял ендпоинт т.к стандартный пакет не может сравнивать схожие ендпоинты в разных роутерах, что приводит к неверному поведению
	apiRouter.Handle("/post/", authMiddlewareManager.AuthMiddleware(posterRouter))
	apiRouter.Handle("/auth/", authRouter)

	http.DefaultServeMux.Handle("/api/", http.StripPrefix("/api", apiRouter))
	server := &http.Server{
		Addr: fmt.Sprintf("%s:%d", cfg.Address, cfg.Port),
	}

	fmt.Println(server.Addr)

	return &HttpServer{
		cfg,
		server,
	}
}

func (s *HttpServer) Start() error {
	return s.http.ListenAndServe()
}

func (s *HttpServer) Stop() error {
	return s.http.Close()
}
