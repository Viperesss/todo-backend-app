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

type validateble interface {
	Validate() error
}

func DecodeAndValidateRequest(r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return fmt.Errorf("decode json: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	var err error

	v, ok := dest.(validateble)
	if ok {
		err = v.Validate()
	} else {
		err = requestValidator.Struct(dest)
	}

	if err != nil {
		err = fmt.Errorf("request validation: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	return err
}
