package task

import (
	"log/slog"

	"github.com/Aibekabdi/time-tracker/internal/service"
)

type Handler struct {
	log      *slog.Logger
	taskRepo service.TaskService
}

func New(log *slog.Logger, taskRepo service.TaskService) *Handler {
	return &Handler{
		log:      log,
		taskRepo: taskRepo,
	}
}
