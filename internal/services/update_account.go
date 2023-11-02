package service

import (
	"context"
	"time"

	model "github.com/gamepkw/accounts-banking-microservice/internal/models"
)

func (a *accountService) UpdateAccount(c context.Context, ar *model.Account) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	*ar.UpdatedAt = time.Now()
	return a.accountRepo.UpdateAccount(ctx, ar)
}
