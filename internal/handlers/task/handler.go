package task

import "log/slog"

type Handler struct {
	log *slog.Logger
}

func New(log *slog.Logger) *Handler {
	return &Handler{
		log: log,
	}
}
