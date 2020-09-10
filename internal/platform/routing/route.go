package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/oudrag/server/internal/platform/application"
)

type Route struct {
	method   Method
	handlers []Handler
}

func NewRoute(method Method) *Route {
	return &Route{method: method}
}

func (r *Route) HandleWith(handlers ...Handler) *Route {
	r.handlers = handlers

	return r
}

func (r *Route) Handlers(c application.Container) ([]gin.HandlerFunc, error) {
	handlers := make([]gin.HandlerFunc, len(r.handlers))
	for i, action := range r.handlers {
		if needInit, ok := action.(application.HasInit); ok {
			if err := needInit.Init(c); err != nil {
				return nil, err
			}
		}

		handlers[i] = action.Handle
	}

	return handlers, nil
}

func (r *Route) Method() string {
	return r.method.String()
}
