package main

import (
	"log"

	"github.com/oudrag/server/cmd/providers"
	"github.com/oudrag/server/internal/platform/application"
)

func main() {
	app := application.NewIoC([]application.ServiceProvider{
		new(providers.GraphServiceProvider),
	})

	log.Fatal(app.Boot())
}
