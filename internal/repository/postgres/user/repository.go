package user

import (
	"context"

	"github.com/Aibekabdi/time-tracker/internal/models"
	def "github.com/Aibekabdi/time-tracker/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ def.UserRepository = (*repository)(nil)

type repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *repository {
	return &repository{
		pool: pool,
	}
}

func (r *repository) Create(ctx context.Context, user models.User) (uint, error)

func (r *repository) Update(ctx context.Context, user models.User) error

func (r *repository) Delete(ctx context.Context, userID uint) error

func (r *repository) Get(ctx context.Context, userID uint) (models.User, error)

func (r *repository) GetALL(ctx context.Context) ([]models.User, error)
