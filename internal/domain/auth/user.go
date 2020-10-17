package auth

import (
	"fmt"
	"time"
)

type RegisterType uint8

var registerTypeValues = []string{"normal", "google_sso"}

func (r RegisterType) Value() string {
	return registerTypeValues[r]
}

func RegisterTypeFromValue(s string) RegisterType {
	for i := 0; i < len(registerTypeValues); i++ {
		if registerTypeValues[i] == s {
			return RegisterType(i)
		}
	}

	return NormalRegistration
}

const (
	NormalRegistration RegisterType = iota
	GoogleSSORegistration
)

type User struct {
	ID           string
	Name         string
	Email        string
	Password     string
	DisplayName  string
	Avatar       string
	RegisteredAt time.Time
	RegisterType RegisterType
}

func (u *User) HasEssentialFields() bool {
	return u.Name != "" && u.Email != ""
}

func (u *User) SetID(id string) {
	u.ID = id
}

func NewUserFromPayload(name, email, avatar string, registerType RegisterType) (*User, error) {
	user := &User{
		Name:         name,
		Email:        email,
		Avatar:       avatar,
		RegisterType: registerType,
	}

	if !user.HasEssentialFields() {
		return nil, fmt.Errorf("payload has not enough data about user")
	}

	return user, nil
}
