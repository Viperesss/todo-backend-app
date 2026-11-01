package users_transport_http

type UsersHttpHandler struct {
	usersService UsersService
}

type UsersService interface {
}

type UsersRepository interface {
}

func newUsersHandler(usersService UsersService) *UsersHttpHandler {
	return &UsersHttpHandler{
		usersService: usersService,
	}
}
