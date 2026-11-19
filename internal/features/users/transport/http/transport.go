package users_transport_http

import (
	"context"
	"net/http"

	"github.com/Viperesss/todo-backend-app/internal/core/domain"
	core_http_server "github.com/Viperesss/todo-backend-app/internal/core/transport/http/server"
)

type UsersService interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUsers(ctx context.Context, limit, offset *int) ([]domain.User, error)
	GetUser(ctx context.Context, id int) (domain.User, error)
}

// Описывает маршруты своей функциональности
type HttpHandler struct {
	usersService UsersService
}

func NewUsersHandler(usersService UsersService) *HttpHandler {
	return &HttpHandler{
		usersService: usersService,
	}
}

func (h *HttpHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: h.GetUsersHandler,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/{id}",
			Handler: h.GetUser,
		},
	}
}
