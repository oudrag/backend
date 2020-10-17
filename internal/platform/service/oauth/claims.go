package oauth

import "github.com/dgrijalva/jwt-go"

type Claim struct {
	jwt.StandardClaims
	Email          string `json:"usr,omitempty"`
	RefreshableTil int64  `json:"rfs,omitempty"`
}
