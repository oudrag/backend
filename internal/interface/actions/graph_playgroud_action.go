package actions

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

type GraphPlaygroundAction struct{}

func (a *GraphPlaygroundAction) Handle(ctx *gin.Context) {
	playground.Handler("GraphQL playground", "/query").
		ServeHTTP(ctx.Writer, ctx.Request)
}
