package users_postgres_repository

import "github.com/Viperesss/todo-backend-app/internal/core/domain"

type UserModel struct {
	ID      int
	Version int

	FullName    string
	PhoneNumber *string
}

func userDomainsFromModels(models []UserModel) []domain.User {
	usersDomain := make([]domain.User, len(models))

	for i, model := range models {
		usersDomain[i] = domain.NewUser(
			model.ID,
			model.Version,
			model.FullName,
			model.PhoneNumber,
		)
	}

	return usersDomain
}

func userDomainFromModel(model UserModel) domain.User {
	return domain.NewUser(
		model.ID,
		model.Version,
		model.FullName,
		model.PhoneNumber,
	)
}
