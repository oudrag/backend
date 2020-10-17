package providers

import (
	"github.com/oudrag/server/internal/core/app"
	"github.com/oudrag/server/internal/domain/auth"
	"github.com/oudrag/server/internal/domain/events"
	"github.com/oudrag/server/internal/platform/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type IndexableRepo interface {
	Index() error
}

var indexableRepos []IndexableRepo

type RepositoryServiceProvider struct{}

func (s RepositoryServiceProvider) Boot(_ app.Container) error {
	for _, repo := range indexableRepos {
		if err := repo.Index(); err != nil {
			return err
		}
	}

	return nil
}

func (s RepositoryServiceProvider) Register(binder app.Binder) {
	binder.Singleton(auth.UserRepositoryBinding, registerUserRepository)
	binder.Singleton(events.EventRepositoryBinding, registerEventRepository)
}

func registerEventRepository(c app.Container) (interface{}, error) {
	var db *mongo.Database
	if err := c.MakeInto(app.DBConnectionBinding, &db); err != nil {
		return nil, err
	}

	eventRepository := repository.NewEventRepository(db)

	return eventRepository, nil
}

func registerUserRepository(c app.Container) (interface{}, error) {
	var db *mongo.Database
	if err := c.MakeInto(app.DBConnectionBinding, &db); err != nil {
		return nil, err
	}

	userRepository := repository.NewUserRepository(db)
	indexableRepos = append(indexableRepos, userRepository)

	return userRepository, nil
}
