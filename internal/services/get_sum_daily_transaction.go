package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
)

func (a *accountService) GetSumDailyTransaction(c context.Context, account_no string) (float64, error) {
	cacheKey := fmt.Sprintf("daily_transaction_%s", account_no)
	sumDailyTransaction, err := a.redis.Get(cacheKey).Result()
	if err != nil {
		return 0, err
	}

	if err == redis.Nil {
		// log.Printf("key: %s not found in redis", cacheKey)
		return 0, nil
	}

	floatSumDailyTransaction, _ := strconv.ParseFloat(sumDailyTransaction, 64)

	return floatSumDailyTransaction, nil
}
