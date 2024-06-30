package user

import (
	"context"

	"github.com/Aibekabdi/time-tracker/internal/models"
	"github.com/Aibekabdi/time-tracker/internal/repository"
	def "github.com/Aibekabdi/time-tracker/internal/service"
)

var _ def.TaskService = (*service)(nil)

type service struct {
	taskRepo   repository.TaskRepository
	apiAddress string
}

func NewService(taskRepo repository.TaskRepository) *service {
	return &service{
		taskRepo: taskRepo,
	}
}

func (s *service) GetALLByUserID(ctx context.Context, userID uint) ([]models.Task, error)

func (s *service) Get(ctx context.Context, taskID uint) (models.Task, error)

func (s *service) Update(ctx context.Context, task models.Task) error

func (s *service) Delete(ctx context.Context, taskID uint) error
