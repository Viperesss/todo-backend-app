package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/Viperesss/todo-backend-app/internal/core/errors"
)

type User struct {
	ID      int
	Version int

	FullName    string
	PhoneNumber *string
}

func NewUser(
	id int,
	version int,
	fullName string,
	PhoneNumber *string,
) User {
	return User{
		ID:          id,
		Version:     version,
		FullName:    fullName,
		PhoneNumber: PhoneNumber,
	}
}

func NewUserUninitialized(fullName string, phoneNumber *string) User {
	return NewUser(
		UnititializedId,
		UnititializedVersion,
		fullName,
		phoneNumber,
	)
}

func (u *User) Validate() error {
	fullNameLen := len([]rune(u.FullName))

	if fullNameLen < 3 || fullNameLen > 100 {
		return fmt.Errorf("invalid `FullName` len: %d: %w", fullNameLen, core_errors.ErrInvalidArgument)
	}

	if u.PhoneNumber != nil {
		phoneNumberLen := len([]rune(*u.PhoneNumber))
		if phoneNumberLen < 10 || phoneNumberLen > 15 {
			return fmt.Errorf("invalid `PhoneNumber` len: %d: %w", phoneNumberLen, core_errors.ErrInvalidArgument)
		}

		// регулярное выражение
		re := regexp.MustCompile(`^\+[0-9]+$`)
		if !re.MatchString(*u.PhoneNumber) {
			return fmt.Errorf("invalid `PhoneNumber` format: %w", core_errors.ErrInvalidArgument)
		}
	}

	return nil
}
