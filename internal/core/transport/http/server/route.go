package core_http_server

import "net/http"

// чтобы каждая отдеальная фича говорила mux сервара о своих требованиях
// Route - описание одного маршрута
type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func NewRoute(method, path string, handler http.HandlerFunc) Route {
	return Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
}
