package actions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oudrag/server/internal/core/app"
	"github.com/oudrag/server/internal/interface/response"
	"golang.org/x/oauth2"
)

type GetAuthSSOURLAction struct {
	googleOAuth *oauth2.Config
}

func (a *GetAuthSSOURLAction) Init(c app.Container) error {
	return c.MakeInto(app.GoogleOAuthBinding, &a.googleOAuth)
}

func (a *GetAuthSSOURLAction) Handle(ctx *gin.Context) {
	service := ctx.DefaultQuery("service", "google")
	v, ok := ctx.Get("token")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.JSON{
			Message: "Something went wrong!",
		})
		return
	}

	state, ok := v.(string)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.JSON{
			Message: "Something went wrong!",
		})
		return
	}

	switch service {
	case "google":
		url := a.googleOAuth.AuthCodeURL(state)

		ctx.JSON(http.StatusOK, response.JSON{
			Data: map[string]interface{}{
				"url": url,
			},
		})
	default:
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.JSON{
			Message: "This Single sign on method is not implemented",
		})
	}
}
