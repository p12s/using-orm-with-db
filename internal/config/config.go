package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

// Config
type Config struct {
	Db      Db
	Backend Backend
}

// Db - current service general db
type Db struct {
	Driver   string `envconfig:"DB_DRIVER" required:"true"`
	Host     string `envconfig:"DB_HOST" required:"true"`
	Port     int    `envconfig:"DB_PORT" required:"true"`
	SslMode  string `envconfig:"DB_SSLMODE" required:"true"`
	User     string `envconfig:"DB_USER" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true"`
	Name     string `envconfig:"DB_NAME" required:"true"`
}

// Backend - current service
type Backend struct {
	Protocol string `envconfig:"BACKEND_PROTOCOL" required:"true"`
	Host     string `envconfig:"BACKEND_HOST" required:"true"`
	Port     int    `envconfig:"BACKEND_PORT" required:"true"`
}

// New - contructor
func New() (*Config, error) {
	cfg := new(Config)

	if err := envconfig.Process("db", &cfg.Db); err != nil {
		return nil, fmt.Errorf("can't load env: %w", err)
	}

	if err := envconfig.Process("backend", &cfg.Backend); err != nil {
		return nil, fmt.Errorf("can't load env: %w", err)
	}

	return cfg, nil
}
