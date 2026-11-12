package core_http_request

import (
	"encoding/json"
	"fmt"
	"net/http"

	core_errors "github.com/Viperesss/todo-backend-app/internal/core/errors"
	"github.com/go-playground/validator/v10"
)

// В общей области видимиости из за того что валидатор использует кэш
var requestValidator = validator.New()

func DecodeAndValidateRequest(r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return fmt.Errorf("decode json: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	if err := requestValidator.Struct(dest); err != nil {
		return fmt.Errorf("request validation: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	return nil
}
