package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/oudrag/server/internal/core/app"
)

type Route struct {
	path     string
	method   Method
	handlers []Handler
}

func Get(path string) *Route {
	return &Route{
		path:   path,
		method: getMethod,
	}
}

func Post(path string) *Route {
	return &Route{
		path:   path,
		method: postMethod,
	}
}

func Put(path string) *Route {
	return &Route{
		path:   path,
		method: putMethod,
	}
}

func Patch(path string) *Route {
	return &Route{
		path:   path,
		method: patchMethod,
	}
}

func Delete(path string) *Route {
	return &Route{
		path:   path,
		method: deleteMethod,
	}
}

func Head(path string) *Route {
	return &Route{
		path:   path,
		method: headMethod,
	}
}

func Options(path string) *Route {
	return &Route{
		path:   path,
		method: optionsMethod,
	}
}

func (r *Route) HandleWith(handlers ...Handler) *Route {
	r.handlers = handlers

	return r
}

func (r *Route) Handlers(c app.Container) ([]gin.HandlerFunc, error) {
	handlers := make([]gin.HandlerFunc, len(r.handlers))
	for i, action := range r.handlers {
		if needInit, ok := action.(app.HasInit); ok {
			if err := needInit.Init(c); err != nil {
				return nil, err
			}
		}

		handlers[i] = action.Handle
	}

	return handlers, nil
}

func (r *Route) Path() string {
	return r.path
}

func (r *Route) Method() string {
	return r.method.String()
}
