package providers

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/oudrag/server/internal/interface/resolvers"
	"github.com/oudrag/server/internal/platform/application"
	"github.com/oudrag/server/internal/platform/gqlcore"
)

type GraphServiceProvider struct{}

func (s GraphServiceProvider) Boot(container application.Container) error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	var resolver *resolvers.Resolver

	if err := container.MakeInto("graph.resolver", &resolver); err != nil {
		return err
	}

	srv := handler.NewDefaultServer(gqlcore.NewExecutableSchema(gqlcore.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	return http.ListenAndServe(":"+port, nil)
}

func (s GraphServiceProvider) Register(binder application.Binder) {
	binder.Bind("graph.resolver", func(app application.Container) (interface{}, error) {
		r := resolvers.NewResolver(app)
		return r, nil
	})
}
