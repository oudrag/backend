package providers

import (
	"fmt"

	"github.com/oudrag/server/internal/core/app"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseServiceProvider struct{}

func (d DatabaseServiceProvider) Register(binder app.Binder) {
	binder.Singleton(app.DBConnectionBinding, func(c app.Container) (interface{}, error) {
		return mongo.NewClient(options.Client().ApplyURI(createDSN()))
	})
}

func createDSN() string {
	host := app.GetEnv(app.DBHost)
	port := app.GetEnv(app.DBPort)
	database := app.GetEnv(app.DBDatabase)
	username := app.GetEnv(app.DBUsername)
	password := app.GetEnv(app.DBPassword)

	userPass := ""
	if username != "" && password != "" {
		userPass = fmt.Sprintf("%s:%s@", username, password)
	}

	return fmt.Sprintf("mongodb://%s%s:%s/%s", userPass, host, port, database)
}
