package service

import (
	"context"

	model "github.com/gamepkw/accounts-banking-microservice/internal/models"
)

func (a *accountService) GetAllAccount(c context.Context, cursor string, num int64) (res []model.Account, nextCursor string, err error) {
	if num == 0 {
		num = 10
	}

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, nextCursor, err = a.accountRepo.GetAllAccount(ctx, cursor, num)
	if err != nil {
		return nil, "", err
	}

	return
}
