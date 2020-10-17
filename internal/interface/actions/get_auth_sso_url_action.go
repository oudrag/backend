package actions

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oudrag/server/internal/core/app"
	"github.com/oudrag/server/internal/interface/response"
	"github.com/oudrag/server/internal/platform/service/oauth"
	"golang.org/x/oauth2"
)

type GetAuthSSOURLAction struct {
	auth *oauth.AuthManager
}

func (a *GetAuthSSOURLAction) Init(c app.Container) error {
	var googleOAuth *oauth2.Config
	if err := c.MakeInto(app.GoogleOAuthBinding, &googleOAuth); err != nil {
		return err
	}

	a.auth = oauth.NewAuthManager(googleOAuth)
	return nil
}

func (a *GetAuthSSOURLAction) Handle(ctx *gin.Context) {
	svc := ctx.DefaultQuery("service", "google")
	state, ok := getToken(ctx)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.JSON{
			Message: "Something went wrong!",
		})
		return
	}

	url := a.auth.GetAuthenticationURL(state, svc)
	if url == "" {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.JSON{
			Message: "Invalid service",
			Errors:  fmt.Sprintf("%s is not a valid SSO service", svc),
		})

		return
	}

	ctx.JSON(http.StatusOK, response.JSON{
		Data: map[string]interface{}{
			"url": url,
		},
	})
}
