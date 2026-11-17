package core_http_utils

import (
	"fmt"
	"net/http"
	"strconv"

	core_errors "github.com/Viperesss/todo-backend-app/internal/core/errors"
)

func GetIntPathValues(r *http.Request, key string) (int, error) {
	pathValue := r.PathValue(key)

	val, err := strconv.Atoi(pathValue)
	if err != nil {
		return -1, fmt.Errorf(
			"path value=`%s`, by key=`%s` not a valid integer: %v: %w",
			pathValue, key, err, core_errors.ErrInvalidArgument,
		)
	}

	return val, nil
}
