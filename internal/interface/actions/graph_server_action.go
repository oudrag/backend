package actions

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/oudrag/server/internal/platform/app"
)

type GraphServerAction struct {
	srv *handler.Server
}

func (a *GraphServerAction) Init(c app.Container) error {
	return c.MakeInto(app.GraphServerBinding, &a.srv)
}

func (a *GraphServerAction) Handle(ctx *gin.Context) {
	a.srv.ServeHTTP(ctx.Writer, ctx.Request)
}
