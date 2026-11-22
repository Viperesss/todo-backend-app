package users_transport_http

import (
	"net/http"

	core_logger "github.com/Viperesss/todo-backend-app/internal/core/logger"
	core_http_response "github.com/Viperesss/todo-backend-app/internal/core/transport/http/response"
	core_http_utils "github.com/Viperesss/todo-backend-app/internal/core/transport/http/utils"
)

func (h *HttpHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPresponseHandler(log, w)

	log.Debug("invoke DeleteUser handler")

	userId, err := core_http_utils.GetIntPathValues(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "couldn't get user id from path value")
		return
	}

	if err := h.usersService.DeleteUser(ctx, userId); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete user")
		return
	}

	responseHandler.NoContentResponse(http.StatusNoContent)
}
