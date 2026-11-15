package users_service

import (
	"context"

	"github.com/Viperesss/todo-backend-app/internal/core/domain"
)

type UsersService struct {
	UsersRepository UsersRepository
}

type UsersRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUsers(ctx context.Context, limit, offset *int) ([]domain.User, error)
}

func NewUsersService(usersRepos UsersRepository) *UsersService {
	return &UsersService{
		UsersRepository: usersRepos,
	}
}
