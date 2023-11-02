package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
)

func (a *accountService) GetDailyLimit(c context.Context, account_no string) (float64, error) {
	cacheKey := fmt.Sprintf("daily_limit_%s", account_no)
	dailyLimit, err := a.redis.Get(cacheKey).Result()
	if err != nil {
		// log.Printf("key: %s not found in redis", cacheKey)
	}

	if err == redis.Nil {
		cacheKey := "default_daily_limit"
		dailyLimit, err := a.redis.Get(cacheKey).Result()
		if err != nil {
			// log.Printf("key: %s not found in redis", cacheKey)
		}
		floatDailyLimit, _ := strconv.ParseFloat(dailyLimit, 64)
		return floatDailyLimit, nil
	}

	floatDailyLimit, _ := strconv.ParseFloat(dailyLimit, 64)

	return floatDailyLimit, nil
}
