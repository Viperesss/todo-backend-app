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
}

type UsersTransport interface {
}

func NewUsersService(usersRepos UsersRepository) *UsersService {
	return &UsersService{
		UsersRepository: usersRepos,
	}
}
