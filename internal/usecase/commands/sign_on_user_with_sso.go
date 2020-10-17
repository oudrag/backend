package commands

import (
	"github.com/oudrag/server/internal/core/app"
	"github.com/oudrag/server/internal/core/cqrs"
	"github.com/oudrag/server/internal/domain/auth"
)

type SignOnUserWithSSOCommand struct {
	Email  string
	Name   string
	Avatar string
}

func (c SignOnUserWithSSOCommand) GetCommandName() string {
	return "SignOnUserWithSSOCommand"
}

type SignOnUserWithSSOCommandHandler struct {
	repo auth.UserRepository
}

func (h *SignOnUserWithSSOCommandHandler) Init(c app.Container) error {
	return c.MakeInto(auth.UserRepositoryBinding, &h.repo)
}

func (h *SignOnUserWithSSOCommandHandler) Handle(cmd *cqrs.Message) cqrs.Response {
	var payload *SignOnUserWithSSOCommand
	if err := cmd.PayloadAs(&payload); err != nil {
		return cqrs.ErrResponse(err)
	}

	var user *auth.User
	user, _ = h.repo.LoadWithEmail(payload.Email)
	if user == nil {
		var err error
		user, err = auth.NewUserFromPayload(
			payload.Name, payload.Email, payload.Avatar,
			auth.GoogleSSORegistration,
		)
		if err != nil {
			return cqrs.ErrResponse(err)
		}

		if err = h.repo.Save(user); err != nil {
			return cqrs.ErrResponse(err)
		}
	}

	return cqrs.OkResponse(user)
}
