package user

import (
	"context"

	"github.com/Aibekabdi/time-tracker/internal/models"
	def "github.com/Aibekabdi/time-tracker/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ def.TaskRepository = (*repository)(nil)

type repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *repository {
	return &repository{
		pool: pool,
	}
}

func (r *repository) GetALLByUserID(ctx context.Context, userID uint) ([]models.Task, error)

func (r *repository) Get(ctx context.Context, taskID uint) (models.Task, error)

func (r *repository) Update(ctx context.Context, task models.Task) error

func (r *repository) Delete(ctx context.Context, taskID uint) error
