package oauth

type googleUserData struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Avatar        string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Locale        string `json:"locale"`
}

func (g googleUserData) GetEmail() string {
	return g.Email
}

func (g googleUserData) GetName() string {
	return g.Name
}

func (g googleUserData) GetAvatarURL() string {
	return g.Avatar
}
