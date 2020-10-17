package auth

const (
	UserRepositoryBinding = "repository.user"
)

type UserData interface {
	GetEmail() string
	GetName() string
	GetAvatarURL() string
}

type UserRepository interface {
	Load(id string) (*User, error)
	LoadWithEmail(email string) (*User, error)
	Save(user *User) error
}
