package users_transport_http

import (
	"encoding/json"
	"net/http"

	errordto "github.com/Viperesss/todo-backend-app/internal/core/errors"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

type CreateUserResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func (h *UsersHttpHandler) CreateUser(w *http.ResponseWriter, r *http.Request) {
	var requestDTO CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&requestDTO); err != nil {
		errDTO := errordto.NewErrDTO(err)

		http.Error(*w, errDTO.Time.GoString(), http.StatusBadRequest)
	}

}
