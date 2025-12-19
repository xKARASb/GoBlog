package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/xkarasb/blog/internal/core/servers"
	"github.com/xkarasb/blog/pkg/db/postgres"
	"github.com/xkarasb/blog/pkg/storage/minio"
)

type Config struct {
	servers.HttpServerConfig
	postgres.PostgresConfig
	minio.MinIOConfig
}

func NewConfig() (*Config, error) {
	cfg := Config{}
	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
