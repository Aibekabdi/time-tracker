package user

import "log/slog"

type Handler struct {
	log        *slog.Logger
	apiAddress string
}

func New(log *slog.Logger, apiAddress string) *Handler {
	return &Handler{
		log: log,
		apiAddress: apiAddress,
	}
}

