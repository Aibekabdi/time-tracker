package user

import (
	"context"

	"github.com/Aibekabdi/time-tracker/internal/models"
	"github.com/Aibekabdi/time-tracker/internal/repository"
	def "github.com/Aibekabdi/time-tracker/internal/service"
)

var _ def.UserService = (*service)(nil)

type service struct {
	userRepo   repository.UserRepository
	apiAddress string
}

func NewService(userRepo repository.UserRepository, apiAddress string) *service {
	return &service{
		userRepo:   userRepo,
		apiAddress: apiAddress,
	}
}

func (s *service) Create(ctx context.Context, user models.User) (uint, error)

func (s *service) Update(ctx context.Context, user models.User) error

func (s *service) Delete(ctx context.Context, userID uint) error

func (s *service) Get(ctx context.Context, userID uint) (models.User, error)

func (s *service) GetALL(ctx context.Context) ([]models.User, error)
