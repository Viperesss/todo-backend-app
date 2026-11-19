package users_transport_http

import (
	"net/http"

	core_http_server "github.com/Viperesss/todo-backend-app/internal/core/transport/http/server"
	users_service "github.com/Viperesss/todo-backend-app/internal/features/users/service"
)

// Описывает маршруты своей функциональности
type HttpHandler struct {
	usersService users_service.UsersService
}

func NewUsersHandler(usersService users_service.UsersService) *HttpHandler {
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
