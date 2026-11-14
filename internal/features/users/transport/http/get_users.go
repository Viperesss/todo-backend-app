package users_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/Viperesss/todo-backend-app/internal/core/logger"
	core_http_response "github.com/Viperesss/todo-backend-app/internal/core/transport/http/response"
	core_http_utils "github.com/Viperesss/todo-backend-app/internal/core/transport/http/utils"
)

type GetUsersResponse []UserDTOResponse

func (h *UsersHttpHandler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPresponseHandler(log, w)

	log.Debug("invoke GetUsers handler")

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "couldn't get limit or offset from query parameters")
		return
	}

	userDomains, err := h.usersService.GetUsers(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get users")
		return
	}

	response := GetUsersResponse(usersDTOFromDomains(userDomains))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	limit, err := core_http_utils.GetIntQueryParam(r, "limit")
	if err != nil {
		return nil, nil, fmt.Errorf("get 'limit' query param: %w", err)
	}

	offset, err := core_http_utils.GetIntQueryParam(r, "offset")
	if err != nil {
		return nil, nil, fmt.Errorf("get 'param' query param: %w", err)
	}

	return limit, offset, nil
}
