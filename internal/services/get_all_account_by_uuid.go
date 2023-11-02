package service

import (
	"context"

	model "github.com/gamepkw/accounts-banking-microservice/internal/models"
)

func (a *accountService) GetAllAccountByUuid(c context.Context, uuid string) (res *[]model.Account, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.accountRepo.GetAllAccountByUuid(ctx, uuid)
	if err != nil {
		return
	}

	return
}
