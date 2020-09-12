package providers

import (
	"github.com/go-redis/redis/v7"
	"github.com/oudrag/server/internal/core/app"
)

type RedisServiceProvider struct{}

func (s *RedisServiceProvider) Register(binder app.Binder) {
	binder.Singleton(app.RedisClientBinding, registerRedisClient)
}

func registerRedisClient(_ app.Container) (interface{}, error) {
	addr := app.GetEnv(app.RedisHost)
	username := app.GetEnv(app.RedisUsername)
	password := app.GetEnv(app.RedisPassword)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: username,
		Password: password,
		DB:       0,
	})

	_, err := client.Ping().Result()

	return client, err
}
