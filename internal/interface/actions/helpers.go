package actions

import (
	"github.com/gin-gonic/gin"
)

func getToken(ctx *gin.Context) (string, bool) {
	v, ok := ctx.Get("token")
	if !ok {
		return "", false
	}

	state, ok := v.(string)

	return state, ok
}
