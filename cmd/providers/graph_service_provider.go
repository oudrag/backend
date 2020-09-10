package providers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/oudrag/server/internal/interface/resolvers"
	"github.com/oudrag/server/internal/platform/application"
	"github.com/oudrag/server/internal/platform/gqlcore"
)

type GraphServiceProvider struct{}

func (s GraphServiceProvider) Register(binder application.Binder) {
	binder.Singleton(gqlcore.ServerBinding, func(app application.Container) (interface{}, error) {
		r := resolvers.NewResolver(app)
		srv := handler.NewDefaultServer(gqlcore.NewExecutableSchema(gqlcore.Config{Resolvers: r}))

		return srv, nil
	})
}
