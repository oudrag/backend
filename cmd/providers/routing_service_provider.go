package providers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/oudrag/server/internal/core/app"
	"github.com/oudrag/server/internal/core/routing"
	"github.com/oudrag/server/internal/interface/actions"
	"github.com/oudrag/server/internal/interface/middleware"
)

type RoutingServiceProvider struct{}

func (s RoutingServiceProvider) Boot(c app.Container) error {
	var router *gin.Engine
	if err := c.MakeInto(app.RouterBinding, &router); err != nil {
		return err
	}

	var routes []*routing.Route
	if err := c.MakeInto(app.RoutesListBinding, &routes); err != nil {
		return err
	}

	if len(routes) == 0 {
		return fmt.Errorf("not any routed defined")
	}

	var middlewares []routing.Handler
	if err := c.MakeInto(app.MiddlewareListBinding, &middlewares); err != nil {
		return err
	}

	// Load global middlewares
	for _, m := range middlewares {
		if needInit, ok := m.(app.HasInit); ok {
			if err := needInit.Init(c); err != nil {
				return err
			}
		}

		router.Use(m.Handle)
	}

	for _, route := range routes {
		handlers, err := route.Handlers(c)
		if err != nil {
			return err
		}

		router.Handle(route.Method(), route.Path(), handlers...)
	}

	return router.Run()
}

func (s RoutingServiceProvider) Register(binder app.Binder) {
	binder.Singleton(app.RouterBinding, registerRouter)
	binder.Singleton(app.RoutesListBinding, registerRoutes)
	binder.Singleton(app.MiddlewareListBinding, registerGlobalMiddlewares)
}

func registerRouter(_ app.Container) (interface{}, error) {
	router := gin.Default()
	router.RedirectTrailingSlash = true

	return router, nil
}

func registerRoutes(_ app.Container) (interface{}, error) {
	return []*routing.Route{
		routing.Get("/").HandleWith(new(actions.GraphPlaygroundAction)),
		routing.Post("/query").HandleWith(new(actions.GraphServerAction)),
		routing.Get("/auth/sso").HandleWith(new(middleware.StateMiddleware), new(actions.GetAuthSSOURLAction)),
		routing.Get("/auth/sso/:service").HandleWith(new(middleware.StateMiddleware), new(actions.GetAuthenticateSSOAction)),
	}, nil
}

func registerGlobalMiddlewares(_ app.Container) (interface{}, error) {
	return []routing.Handler{}, nil
}
