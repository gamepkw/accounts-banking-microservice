package service

import (
	"context"

	model "github.com/gamepkw/accounts-banking-microservice/internal/models"
)

func (a *accountService) GetAccountByAccountNo(c context.Context, account_no string) (res *model.Account, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.accountRepo.GetAccountByAccountNo(ctx, account_no)
	if err != nil {
		return
	}

	return
}
