package actions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oudrag/server/internal/core/app"
	"github.com/oudrag/server/internal/core/cqrs"
	"github.com/oudrag/server/internal/domain/auth"
	"github.com/oudrag/server/internal/interface/response"
	"github.com/oudrag/server/internal/platform/service/oauth"
	"github.com/oudrag/server/internal/usecase/commands"
	"golang.org/x/oauth2"
)

type GetAuthenticateSSOAction struct {
	auth       *oauth.AuthManager
	commandBus *cqrs.CommandBus
}

func (a *GetAuthenticateSSOAction) Init(c app.Container) error {
	var googleOAuth *oauth2.Config
	if err := c.MakeInto(app.CQRSCommandBusBinding, &a.commandBus); err != nil {
		return err
	}
	if err := c.MakeInto(app.GoogleOAuthBinding, &googleOAuth); err != nil {
		return err
	}
	a.auth = oauth.NewAuthManager(googleOAuth)

	return nil
}

func (a *GetAuthenticateSSOAction) Handle(ctx *gin.Context) {
	srv := ctx.Param("service")
	state, ok := getToken(ctx)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.JSON{
			Message: "Something went wrong!",
		})
		return
	}
	if state == ctx.Query("state") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.JSON{
			Message: "Invalid Action",
		})
		return
	}

	userData, err := a.auth.GetUserData(ctx.Query("code"), srv)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.JSON{
			Message: "Invalid Code or Internal Error",
			Errors:  err.Error(),
		})
		return
	}

	res := a.commandBus.Dispatch(
		cqrs.NewCommand(&commands.SignOnUserWithSSOCommand{
			Email:  userData.GetEmail(),
			Name:   userData.GetName(),
			Avatar: userData.GetAvatarURL(),
		}),
	)

	if !res.Ok() {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.JSON{
			Message: "Something went wrong!",
		})
		return
	}

	user, ok := res.Result().(*auth.User)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.JSON{
			Message: "Something went wrong!",
		})
		return
	}

	jwt, err := a.auth.GenerateAccessToken(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.JSON{
			Message: "Something went wrong!",
		})
		return
	}

	ctx.JSON(http.StatusOK, response.JSON{
		Message: "Successfully Authenticated!",
		Data: map[string]interface{}{
			"token": jwt,
			"type":  "bearer",
		},
	})
}
