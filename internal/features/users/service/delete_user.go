package users_service

import "context"

func (s *Service) DeleteUser(ctx context.Context, id int) error {
	return s.usersRepository.DeleteUser(ctx, id)
}
