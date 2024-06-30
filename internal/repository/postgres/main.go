package postgres

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Aibekabdi/time-tracker/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(ctx context.Context, logger *slog.Logger, cfg *config.DB) (*pgxpool.Pool, error) {
	const op = "postgres.Connect"
	log := logger.With(slog.String("op", op))
	
	log.Info("init...")
	pool, err := pgxpool.New(ctx, cfg.DBURL)
	if err != nil {
		return nil, err
	}

	log.Info("checking conn...")
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	conn.Release()
	log.Info("checking conn is succesfull")
	log.Info("init done")

	log.Info("setting time zone", slog.String("time zone", cfg.TimeZone))
	if _, err := pool.Exec(ctx, fmt.Sprintf("SET TIME ZONE '%s'", cfg.TimeZone)); err != nil {
		return nil, err
	}

	return pool, nil
}
