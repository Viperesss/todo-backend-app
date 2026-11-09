package users_transport_http

import (
	"net/http"

	core_logger "github.com/Viperesss/todo-backend-app/internal/core/logger"
	core_http_request "github.com/Viperesss/todo-backend-app/internal/core/transport/http/request"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name" validate:"required,min=3,max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,e164"`
}

type CreateUserResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func (h *UsersHttpHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	log.Debug("invoke CreateUser handler")

	var requestDTO CreateUserRequest

	if err := core_http_request.DecodeAndValidateRequest(r, &requestDTO); err != nil {

	}

	w.WriteHeader(http.StatusOK)
}
