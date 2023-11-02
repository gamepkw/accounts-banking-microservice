package service

import (
	"context"

	model "github.com/gamepkw/accounts-banking-microservice/internal/models"
)

func (a *accountService) RegisterAccount(c context.Context, m *model.Account) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	if err = a.GenerateAccountNo(ctx, m); err != nil {
		return err
	}

	if err = a.accountRepo.RegisterAccount(ctx, m); err != nil {
		return err
	}

	return
}
