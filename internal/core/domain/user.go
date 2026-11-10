package domain

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
