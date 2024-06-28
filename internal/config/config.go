package config

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	EnvLevel   string `envconfig:"ENV_LEVEL" default:"local"`
	ApiAddress string `envconfig:"API_ADDRESS"`
	HTTPServer
	DB
}

type HTTPServer struct {
	Address     string        `envconfig:"SERVER_ADDRESS" default:"127.0.0.1:8000"`
	Timeout     time.Duration `envconfig:"SERVER_TIMEOUT" default:"4s"`
	IdleTimeout time.Duration `envconfig:"SERVER_IDLE_TIMEOUT" default:"30s"`
}

type DB struct {
}

func LoadConfig() (*Config, error) {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		return nil, fmt.Errorf("unvailable to load config: %s", err.Error())
	}
	return &config, nil
}
