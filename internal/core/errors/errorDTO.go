package errordto

import (
	"encoding/json"
	"time"
)

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func NewErrDTO(err error) ErrorDTO {
	return ErrorDTO{
		Message: err.Error(),
		Time:    time.Now(),
	}
}

func (e *ErrorDTO) ToString(errorDTO ErrorDTO) (string, error) {
	b, err := json.Marshal(e)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
