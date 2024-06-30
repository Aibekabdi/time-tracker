package user

import (
	"log/slog"

	"github.com/Aibekabdi/time-tracker/internal/service"
)

type Handler struct {
	log         *slog.Logger
	apiAddress  string
	userService service.UserService
}

func New(log *slog.Logger, apiAddress string, userService service.UserService) *Handler {
	return &Handler{
		log:         log,
		apiAddress:  apiAddress,
		userService: userService,
	}
}
