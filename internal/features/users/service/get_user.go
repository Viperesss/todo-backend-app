package users_service

import (
	"context"
	"fmt"

	"github.com/Viperesss/todo-backend-app/internal/core/domain"
)

func (s *UsersService) GetUser(ctx context.Context, id int) (domain.User, error) {
	if id < 1 {
		return domain.User{}, fmt.Errorf("id must be a positive number")
	}

	user, err := s.UsersRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user: %w", err)
	}

	return user, nil
}
