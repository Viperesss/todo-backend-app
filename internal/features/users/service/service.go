package users_service

import (
	"context"

	"github.com/Viperesss/todo-backend-app/internal/core/domain"
)

type UsersRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUsers(ctx context.Context, limit, offset *int) ([]domain.User, error)
	GetUser(ctx context.Context, id int) (domain.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type Service struct {
	usersRepository UsersRepository
}

func NewService(usersRepos UsersRepository) *Service {
	return &Service{
		usersRepository: usersRepos,
	}
}
