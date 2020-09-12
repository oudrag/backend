package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/oudrag/server/internal/core/app"
	"github.com/oudrag/server/internal/core/helpers"
)

type StateMiddleware struct {
	redis *redis.Client
}

func (m *StateMiddleware) Init(c app.Container) error {
	return c.MakeInto(app.RedisClientBinding, &m.redis)
}

func (m *StateMiddleware) Handle(ctx *gin.Context) {
	fingerprint := ctx.GetHeader("fingerprint")

	token, err := m.redis.Get(fingerprint).Result()
	if err != nil {
		hFingerprint := helpers.HashSHA1(fingerprint)
		hTime := helpers.HashSHA1(time.Now().String())
		token = hFingerprint + "." + hTime

		m.redis.Set(fingerprint, token, time.Minute*5)
	}

	ctx.Set("fingerprint", fingerprint)
	ctx.Set("token", token)
}
