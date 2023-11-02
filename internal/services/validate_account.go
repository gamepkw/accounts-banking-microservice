package service

import (
	"context"

	model "github.com/gamepkw/accounts-banking-microservice/internal/models"
)

func (a *accountService) ValidateAccount(c context.Context, ar *model.Account) (err error) {

	if ar.IsClosed == 1 {
		return model.ErrNotFound
	}
	return
}
