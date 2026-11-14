package users_transport_http

import "github.com/Viperesss/todo-backend-app/internal/core/domain"

type UserDTOResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func userDTOFromDomain(userDomain domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:          userDomain.ID,
		Version:     userDomain.Version,
		FullName:    userDomain.FullName,
		PhoneNumber: userDomain.PhoneNumber,
	}
}

func usersDTOFromDomains(userDomains []domain.User) []UserDTOResponse {
	usersDTO := make([]UserDTOResponse, len(userDomains))

	for i, user := range userDomains {
		usersDTO[i] = userDTOFromDomain(user)
	}

	return usersDTO
}
