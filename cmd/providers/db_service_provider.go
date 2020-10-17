package providers

import (
	"context"
	"fmt"
	"time"

	"github.com/oudrag/server/internal/core/app"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseServiceProvider struct{}

func (s DatabaseServiceProvider) Register(binder app.Binder) {
	binder.Singleton(app.DBConnectionBinding, registerMongoConnection)
}

func registerMongoConnection(_ app.Container) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(createDSN()))
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(app.GetEnv(app.DBDatabase)), nil
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
