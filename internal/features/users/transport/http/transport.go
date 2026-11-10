package users_transport_http

import (
	"context"
	"net/http"

	"github.com/Viperesss/todo-backend-app/internal/core/domain"
	core_http_server "github.com/Viperesss/todo-backend-app/internal/core/transport/http/server"
)

// Описывает маршруты своей функциональности
type UsersHttpHandler struct {
	usersService UsersService
}

type UsersService interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
}

type UsersRepository interface {
}

func NewUsersHandler(usersService UsersService) *UsersHttpHandler {
	return &UsersHttpHandler{
		usersService: usersService,
	}
}

func (h *UsersHttpHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
	}
}
