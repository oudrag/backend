package providers

import (
	"github.com/oudrag/server/internal/core/app"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OAuthServiceProvider struct{}

func (g OAuthServiceProvider) Register(binder app.Binder) {
	binder.Singleton(app.GoogleOAuthBinding, registerGoogleOAuth)
}

func registerGoogleOAuth(_ app.Container) (interface{}, error) {
	cid := app.GetEnv(app.GoogleClientID)
	secret := app.GetEnv(app.GoogleClientSecret)

	cnf := &oauth2.Config{
		ClientID:     cid,
		ClientSecret: secret,
		Endpoint:     google.Endpoint,
		RedirectURL:  app.GetEnv(app.AppURL) + "/auth/sso/google",
		Scopes: []string{
			"profile", "email", "openid",
		},
	}

	return cnf, nil
}
