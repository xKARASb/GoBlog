package main

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/xkarasb/blog/internal/core/servers"
)

func main() {
	cfg := servers.HttpServerConfig{}
	cleanenv.ReadConfig(".env", &cfg)
	serv := servers.NewHttpServer(cfg)

	fmt.Println(serv.Start())
}
