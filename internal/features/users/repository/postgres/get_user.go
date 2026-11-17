package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Viperesss/todo-backend-app/internal/core/domain"
)

func (r *UsersRepository) GetUser(ctx context.Context, id int) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, full_name, phone_number
	FROM todoapp.users
	WHERE id = $1;
	`

	var userModel UserModel

	row := r.pool.QueryRow(ctx, query, id)
	if err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.PhoneNumber,
	); err != nil {
		return domain.User{}, fmt.Errorf("scan row: %w", err)
	}

	userDomain := userDomainFromModel(userModel)

	return userDomain, nil
}
