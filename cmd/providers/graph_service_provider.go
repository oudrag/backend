package providers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/oudrag/server/internal/core/app"
	"github.com/oudrag/server/internal/core/gqlcore"
	"github.com/oudrag/server/internal/interface/resolvers"
)

type GraphServiceProvider struct{}

func (s GraphServiceProvider) Register(binder app.Binder) {
	binder.Singleton(app.GraphServerBinding, func(app app.Container) (interface{}, error) {
		r := resolvers.NewResolver(app)
		srv := handler.NewDefaultServer(gqlcore.NewExecutableSchema(gqlcore.Config{Resolvers: r}))

		return srv, nil
	})
}
