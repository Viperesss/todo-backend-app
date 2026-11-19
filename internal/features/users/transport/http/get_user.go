package users_transport_http

import (
	"net/http"

	core_logger "github.com/Viperesss/todo-backend-app/internal/core/logger"
	core_http_response "github.com/Viperesss/todo-backend-app/internal/core/transport/http/response"
	core_http_utils "github.com/Viperesss/todo-backend-app/internal/core/transport/http/utils"
)

type GetUserResponse UserDTOResponse

func (h *UsersHttpHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPresponseHandler(log, w)

	log.Debug("invoke GetUser handler")

	userID, err := core_http_utils.GetIntPathValues(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "couldn't get user id from path value")
		return
	}

	userDomain, err := h.usersService.GetUser(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user")
		return
	}

	userDTO := GetUserResponse(userDTOFromDomain(userDomain))

	responseHandler.JSONResponse(userDTO, http.StatusOK)
}
