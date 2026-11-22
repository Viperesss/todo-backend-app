package core_http_types

import (
	"encoding/json"
	"fmt"

	"github.com/Viperesss/todo-backend-app/internal/core/domain"
)

type Nullable[T any] struct {
	domain.Nullable[T]
}

// переопределение сиандартного метода json.Decode;
// реализация интерфейса json.Unmarshaler
func (n *Nullable[T]) UnmarshalJSON(b []byte) error {
	// если этот метод был вызван, то 100% поле было передано
	n.Set = true

	if string(b) == "null" {
		n.Value = nil

		return nil
	}

	var value T
	if err := json.Unmarshal(b, &value); err != nil {
		return fmt.Errorf("json unmarshal: %w", err)
	}

	n.Value = &value

	return nil
}

func (n *Nullable[T]) ToDomain() domain.Nullable[T] {
	return domain.Nullable[T]{
		Value: n.Value,
		Set:   n.Set,
	}
}
