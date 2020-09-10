package providers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/oudrag/server/internal/interface/actions"
	"github.com/oudrag/server/internal/platform/application"
	"github.com/oudrag/server/internal/platform/routing"
)

type RoutingServiceProvider struct{}

func (s RoutingServiceProvider) Boot(c application.Container) error {
	var router *gin.Engine
	if err := c.MakeInto(routing.RouterBinding, &router); err != nil {
		return err
	}

	var routes map[string]*routing.Route
	if err := c.MakeInto(routing.RoutesListBinding, &routes); err != nil {
		return err
	}

	if len(routes) == 0 {
		return fmt.Errorf("not any routed defined")
	}

	var middlewares []routing.Handler
	if err := c.MakeInto(routing.MiddlewareListBinding, &middlewares); err != nil {
		return err
	}

	// Load global middlewares
	for _, m := range middlewares {
		if needInit, ok := m.(application.HasInit); ok {
			if err := needInit.Init(c); err != nil {
				return err
			}
		}

		router.Use(m.Handle)
	}

	for path, route := range routes {
		handlers, err := route.Handlers(c)
		if err != nil {
			return err
		}

		router.Handle(route.Method(), path, handlers...)
	}

	return router.Run()
}

func (s RoutingServiceProvider) Register(binder application.Binder) {
	binder.Singleton(routing.RouterBinding, registerRouter)
	binder.Singleton(routing.RoutesListBinding, registerRoutes)
	binder.Singleton(routing.MiddlewareListBinding, registerGlobalMiddlewares)
}

func registerRouter(_ application.Container) (interface{}, error) {
	router := gin.Default()
	router.RedirectTrailingSlash = true

	return router, nil
}

func registerRoutes(_ application.Container) (interface{}, error) {
	return map[string]*routing.Route{
		"/":      routing.NewRoute(routing.Get).HandleWith(new(actions.GraphPlaygroundAction)),
		"/query": routing.NewRoute(routing.Post).HandleWith(new(actions.GraphServerAction)),
	}, nil
}

func registerGlobalMiddlewares(_ application.Container) (interface{}, error) {
	return []routing.Handler{}, nil
}
