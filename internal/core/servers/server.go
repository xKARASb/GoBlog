package servers

import (
	"fmt"
	"net/http"

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
	// Передают родительский роутер, что бы был равный уровень некоторых ендпоинтов
	// т.к стандартный пакет не может сравнивать схожие ендпоинты в разных роутерах, что приводит к неверному поведению
	readRouter := routers.GetReaderRouter(apiRouter)
	posterRouter := routers.GetPosterRouter(apiRouter)
	authRouter := routers.GetAuthRouter(apiRouter, db)

	apiRouter.Handle("/", readRouter)
	apiRouter.Handle("/posts/", posterRouter)
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
