package service

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	model "github.com/gamepkw/accounts-banking-microservice/internal/models"
	transactionModel "github.com/gamepkw/transactions-banking-microservice/models"

	"github.com/go-redis/redis"
)

type accountService struct {
	accountRepo     model.AccountRepository
	transactionRepo transactionModel.TransactionRepository
	contextTimeout  time.Duration
	redis           *redis.Client
}

// NewAccountService will create new an accountService object representation of model.AccountService interface
func NewAccountService(ar model.AccountRepository, tr transactionModel.TransactionRepository, redis *redis.Client, timeout time.Duration) model.AccountService {
	return &accountService{
		accountRepo:     ar,
		transactionRepo: tr,
		redis:           redis,
		contextTimeout:  timeout,
	}
}

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

func (a *accountService) GetAllAccountByUuid(c context.Context, uuid string) (res *[]model.Account, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.accountRepo.GetAllAccountByUuid(ctx, uuid)
	if err != nil {
		return
	}

	return
}

func (a *accountService) GetAccountByAccountNo(c context.Context, account_no string) (res *model.Account, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.accountRepo.GetAccountByAccountNo(ctx, account_no)
	if err != nil {
		return
	}

	return
}

func (a *accountService) GetCountAccount(c context.Context) (res map[string]int, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.accountRepo.GetCountAccountByStatus(ctx)
	if err != nil {
		return
	}

	return
}

func (a *accountService) UpdateAccount(c context.Context, ar *model.Account) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	*ar.UpdatedAt = time.Now()
	return a.accountRepo.UpdateAccount(ctx, ar)
}

func (a *accountService) RegisterAccount(c context.Context, m *model.Account) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	if err = a.GenerateAccountNo(ctx, m); err != nil {
		return err
	}

	if err = a.accountRepo.RegisterAccount(ctx, m); err != nil {
		return err
	}

	return
}

func (a *accountService) GenerateAccountNo(c context.Context, m *model.Account) (err error) {
	rand.Seed(time.Now().UnixNano())
	digit_1 := rand.Intn(10)
	digit_2 := rand.Intn(10)
	digit_3 := rand.Intn(10)
	digit_4 := rand.Intn(10)
	digit_5 := rand.Intn(10)
	digit_6 := rand.Intn(10)
	digit_7 := rand.Intn(10)
	digit_8 := rand.Intn(10)
	digit_9 := rand.Intn(10)
	digit_10 := (digit_1 + digit_2 + digit_3 + digit_4 + digit_5 + digit_6 + digit_7 + digit_8 + digit_9) % 10
	str_digit_1 := strconv.Itoa(digit_1)
	str_digit_2 := strconv.Itoa(digit_2)
	str_digit_3 := strconv.Itoa(digit_3)
	str_digit_4 := strconv.Itoa(digit_4)
	str_digit_5 := strconv.Itoa(digit_5)
	str_digit_6 := strconv.Itoa(digit_6)
	str_digit_7 := strconv.Itoa(digit_7)
	str_digit_8 := strconv.Itoa(digit_8)
	str_digit_9 := strconv.Itoa(digit_9)
	str_digit_10 := strconv.Itoa(digit_10)
	m.AccountNo = (str_digit_1 + str_digit_2 + str_digit_3 + str_digit_4 + str_digit_5 + str_digit_6 + str_digit_7 + str_digit_8 + str_digit_9 + str_digit_10)

	return
}

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

func (a *accountService) ValidateAccount(c context.Context, ar *model.Account) (err error) {

	if ar.IsClosed == 1 {
		return model.ErrNotFound
	}
	return
}

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
