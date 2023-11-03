package service

import "context"

func (a *accountService) GetDailyRemainingAmount(c context.Context, accountNo string) (float64, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	sumDailyTransaction, _ := a.GetSumDailyTransaction(ctx, accountNo)

	dailyLimitPerDay, _ := a.GetDailyLimit(ctx, accountNo)

	remainingAmount := dailyLimitPerDay - sumDailyTransaction

	return remainingAmount, nil
}
