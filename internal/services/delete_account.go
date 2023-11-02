package service

import (
	"context"

	model "github.com/gamepkw/accounts-banking-microservice/internal/models"
)

func (a *accountService) DeleteAccount(c context.Context, account_no string) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedAccount, err := a.accountRepo.GetAccountFromRedisByAccountNo(ctx, account_no)
	if err != nil {
		return
	}
	if existedAccount == (&model.Account{}) {
		return model.ErrNotFound
	}
	return a.accountRepo.DeleteAccount(ctx, account_no)
}
