package repository

import (
	"context"

	"github.com/Aibekabdi/time-tracker/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user models.User) (uint, error)
	Update(ctx context.Context, user models.User) error
	Delete(ctx context.Context, userID uint) error
	Get(ctx context.Context, userID uint) (models.User, error)
	GetALL(ctx context.Context) ([]models.User, error) // TODO add pagination
}

type TaskRepository interface {
	GetALLByUserID(ctx context.Context, userID uint) ([]models.Task, error) // TODO add pagination
	Get(ctx context.Context, taskID uint) (models.Task, error)
	Update(ctx context.Context, task models.Task) error
	Delete(ctx context.Context, taskID uint) error
}
