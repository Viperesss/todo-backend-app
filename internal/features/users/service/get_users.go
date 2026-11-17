package users_service

import (
	"context"
	"fmt"

	"github.com/Viperesss/todo-backend-app/internal/core/domain"
	core_errors "github.com/Viperesss/todo-backend-app/internal/core/errors"
)

func (s *UsersService) GetUsers(ctx context.Context, limit, offset *int) ([]domain.User, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf("limit must be non-negatime: %w", core_errors.ErrInvalidArgument)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf("offset must be non-negatime: %w", core_errors.ErrInvalidArgument)
	}

	usersDomain, err := s.UsersRepository.GetUsers(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get users: %w", err)
	}

	return usersDomain, nil
}
