package users_service

import (
	"context"

	"github.com/Viperesss/todo-backend-app/internal/core/domain"
	users_postgres_repository "github.com/Viperesss/todo-backend-app/internal/features/users/repository/postgres"
)

type UsersService interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUsers(ctx context.Context, limit, offset *int) ([]domain.User, error)
	GetUser(ctx context.Context, id int) (domain.User, error)
}

type Service struct {
	UsersRepository users_postgres_repository.UsersRepository
}

func NewUsersService(usersRepos users_postgres_repository.UsersRepository) *Service {
	return &Service{
		UsersRepository: usersRepos,
	}
}
