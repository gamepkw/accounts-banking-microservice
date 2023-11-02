package service

import (
	"context"
)

func (a *accountService) GetCountAccount(c context.Context) (res map[string]int, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.accountRepo.GetCountAccountByStatus(ctx)
	if err != nil {
		return
	}

	return
}
