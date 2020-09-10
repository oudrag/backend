package actions

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/oudrag/server/internal/platform/application"
)

type GraphServerAction struct {
	srv *handler.Server
}

func (a *GraphServerAction) Init(c application.Container) error {
	return c.MakeInto(application.GraphServerBinding, &a.srv)
}

func (a *GraphServerAction) Handle(ctx *gin.Context) {
	a.srv.ServeHTTP(ctx.Writer, ctx.Request)
}
