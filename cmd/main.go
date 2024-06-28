package main

import (
	"errors"
	"log"
	"log/slog"
	"os"
	"time-tracker/internal/config"
)

const (
	envLocal = "local"
)

var (
	errSetupLogger = errors.New("invalid level of application")
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	_, err = setupLogger(conf.EnvLevel)
	if err != nil {
		log.Fatal(err)
	}
}

func setupLogger(env string) (*slog.Logger, error) {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	default:
		return nil, errSetupLogger
	}
	return log, nil
}
