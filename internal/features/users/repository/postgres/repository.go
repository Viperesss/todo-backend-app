package users_postgres_repository

import (
	"context"

	"github.com/Viperesss/todo-backend-app/internal/core/domain"
	core_postgres_pool "github.com/Viperesss/todo-backend-app/internal/core/repository/postgres/pool"
)

type UsersRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUsers(ctx context.Context, limit, offset *int) ([]domain.User, error)
	GetUser(ctx context.Context, id int) (domain.User, error)
}

type Repository struct {
	pool core_postgres_pool.Pool
}

func NewUsersRepository(pool core_postgres_pool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}
