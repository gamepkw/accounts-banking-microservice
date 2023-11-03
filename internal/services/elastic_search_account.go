package service

import (
	"context"

	model "github.com/gamepkw/accounts-banking-microservice/internal/models"
)

func (a *accountService) ElasticSearchAccountByAccountNo(c context.Context, account model.ElasticSearchAccount) (res *[]string, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.accountRepo.ElasticSearchAccountByAccountNo(ctx, account)
	if err != nil {
		return
	}

	return
}
