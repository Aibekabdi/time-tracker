package main

import (
	"errors"
	"log"
	"log/slog"
	"os"

	"github.com/Aibekabdi/time-tracker/internal/app"
	"github.com/Aibekabdi/time-tracker/internal/config"
	"github.com/joho/godotenv"
)

const (
	envLocal = "local"
)

var (
	errSetupLogger = errors.New("invalid level of application")
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	logger, err := setupLogger(conf.EnvLevel)
	if err != nil {
		log.Fatal(err)
	}
	app := app.New(logger, conf)
	app.Run()
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
