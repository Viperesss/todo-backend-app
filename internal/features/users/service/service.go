package user_service

import (
	"context"

	"github.com/Viperesss/todo-backend-app/internal/core/domain"
)

type UsersService struct {
	UsersRepository UsersRepository
	UsersTransport  UsersTransport
}

type UsersRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
}

type UsersTransport interface {
}

func NewUsersService(usersRepos UsersRepository, usersTrans UsersTransport) *UsersService {
	return &UsersService{
		UsersRepository: usersRepos,
		UsersTransport:  usersTrans,
	}
}
