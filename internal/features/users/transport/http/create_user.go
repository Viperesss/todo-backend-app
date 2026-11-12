package users_transport_http

import (
	"net/http"

	"github.com/Viperesss/todo-backend-app/internal/core/domain"
	core_logger "github.com/Viperesss/todo-backend-app/internal/core/logger"
	core_http_request "github.com/Viperesss/todo-backend-app/internal/core/transport/http/request"
	core_http_response "github.com/Viperesss/todo-backend-app/internal/core/transport/http/response"
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
	responseHandler := core_http_response.NewHTTPresponseHandler(log, w)

	log.Debug("invoke CreateUser handler")

	var requestDTO CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &requestDTO); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")

		return
	}

	userDomain := domainFromReqDTO(requestDTO)

	newUserDomain, err := h.usersService.CreateUser(ctx, userDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")
		return
	}

	responseDTO := respDTOFromDomain(newUserDomain)

	responseHandler.JSONResponse(responseDTO, http.StatusCreated)
}

func domainFromReqDTO(user CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(user.FullName, user.PhoneNumber)
}

func respDTOFromDomain(userDomain domain.User) CreateUserResponse {
	return CreateUserResponse{
		ID:          userDomain.ID,
		Version:     userDomain.Version,
		FullName:    userDomain.FullName,
		PhoneNumber: userDomain.PhoneNumber,
	}
}
