package users_postgres_repository

import core_postgres_pool "github.com/Viperesss/todo-backend-app/internal/core/repository/postgres/pool"

type UsersRepository struct {
	pool core_postgres_pool.ConnectionPool

	UsersService UsersService
}

type UsersService interface {
}

func NewUsersRepository(pool *core_postgres_pool.ConnectionPool) *UsersRepository {
	return &UsersRepository{
		pool: *pool,
	}
}
