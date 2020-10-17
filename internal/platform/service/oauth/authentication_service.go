package oauth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/oudrag/server/internal/core/app"
	"github.com/oudrag/server/internal/domain/auth"
	"golang.org/x/oauth2"
)

const (
	google = "google"
)

type AuthManager struct {
	googleOAuth *oauth2.Config
}

func NewAuthManager(googleOAuth *oauth2.Config) *AuthManager {
	return &AuthManager{googleOAuth: googleOAuth}
}

func (a *AuthManager) GetAuthenticationURL(state, service string) string {
	switch service {
	case google:
		return a.googleOAuth.AuthCodeURL(state)
	default:
		return ""
	}
}

func (a *AuthManager) GetUserData(code, service string) (auth.UserData, error) {
	switch service {
	case google:
		return a.GetUserDataViaGoogle(code)
	default:
		return nil, fmt.Errorf("invalid service")
	}
}

func (a *AuthManager) GetUserDataViaGoogle(code string) (auth.UserData, error) {
	token, err := a.googleOAuth.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	client := a.googleOAuth.Client(context.Background(), token)
	data, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return nil, err
	}

	defer data.Body.Close()
	rawData, _ := ioutil.ReadAll(data.Body)

	var userData *googleUserData
	err = json.Unmarshal(rawData, &userData)

	return userData, err
}

func (a *AuthManager) GenerateAccessToken(user *auth.User) (string, error) {
	now := time.Now()
	claims := Claim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.AddDate(0, 0, 1).Unix(),
		},
		Email:          user.Email,
		RefreshableTil: now.AddDate(0, 0, 7).Unix(),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(app.GetEnv(app.JWTSecret))
}
