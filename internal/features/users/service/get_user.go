package users_service

import (
	"context"
	"fmt"

	"github.com/Viperesss/todo-backend-app/internal/core/domain"
)

func (s *Service) GetUser(ctx context.Context, id int) (domain.User, error) {
	user, err := s.UsersRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user from repository: %w", err)
	}

	return user, nil
}
