package core_http_middleware

import (
	"net/http"

	core_logger "github.com/Viperesss/todo-backend-app/internal/core/logger"
)

type Middleware func(http.Handler) http.Handler

func ChainMiddleware(log *core_logger.Logger, h http.Handler, m ...Middleware) http.Handler {
	if len(m) == 0 {
		log.Debug("no middleware provided, returning original handler")
		return h
	}

	for i := len(m) - 1; i >= 0; i-- {
		h = m[i](h)
	}

	return h
}
