package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Aibekabdi/time-tracker/internal/config"
	"github.com/Aibekabdi/time-tracker/internal/handlers"
	"github.com/Aibekabdi/time-tracker/internal/repository/postgres"
	"github.com/Aibekabdi/time-tracker/internal/server"
)

type App struct {
	log *slog.Logger
	cfg *config.Config
}

func New(log *slog.Logger, cfg *config.Config) *App {
	return &App{
		log: log,
		cfg: cfg,
	}
}

func (a *App) Run() {
	const op = "app.Run"
	log := a.log.With(slog.String("op", op))
	ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancel()
	pool, err := postgres.NewPostgres(ctx, a.log, &a.cfg.DB)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	handlers := handlers.NewHandler(a.log, a.cfg.ApiAddress)
	srv := new(server.Server)

	log.Info("server is starting", slog.String("address", a.cfg.HTTPServer.Address))
	go func() {
		if err := srv.Run(a.cfg.HTTPServer, handlers.InitRoutes(pool)); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Error("failed to start server", slog.String("error", err.Error()))
				os.Exit(2)
			}
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	log.Info("Received terminate, graceful shutdown", slog.Any("signal", sig))

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("failed gracefull shutdown", slog.Any("error", err))
		os.Exit(3)
	}
	log.Info("Server is closed")
}
