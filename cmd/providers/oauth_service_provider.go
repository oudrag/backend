package providers

import (
	"github.com/oudrag/server/internal/core/app"
	"github.com/oudrag/server/internal/core/url"
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
		RedirectURL:  url.NewAppUrl().AddURI("/auth/sso/google").String(),
		Scopes: []string{
			"profile", "email", "openid",
		},
	}

	return cnf, nil
}
