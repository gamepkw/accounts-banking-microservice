package service

import (
	"context"
	"time"

	model "github.com/gamepkw/accounts-banking-microservice/internal/models"
	accountRepo "github.com/gamepkw/accounts-banking-microservice/internal/repositories"

	"github.com/go-redis/redis"
)

type accountService struct {
	accountRepo    accountRepo.AccountRepository
	contextTimeout time.Duration
	redis          *redis.Client
}

func NewAccountService(
	ar accountRepo.AccountRepository,
	redis *redis.Client,
	timeout time.Duration) AccountService {
	return &accountService{
		accountRepo:    ar,
		redis:          redis,
		contextTimeout: timeout,
	}
}

type AccountService interface {
	GetAllAccount(ctx context.Context, cursor string, num int64) ([]model.Account, string, error)
	GetAccountByAccountNo(ctx context.Context, account_no string) (*model.Account, error)
	UpdateAccount(ctx context.Context, ar *model.Account) error
	RegisterAccount(context.Context, *model.Account) error
	DeleteAccount(ctx context.Context, account_no string) error
	GetCountAccount(ctx context.Context) (map[string]int, error)
	ValidateAccount(ctx context.Context, ar *model.Account) error
	GetAllAccountByUuid(c context.Context, uuid string) (res *[]model.Account, err error)
	GetDailyLimit(c context.Context, account_no string) (float64, error)
	GetSumDailyTransaction(c context.Context, account_no string) (float64, error)
	SelectBank(lastDigit string) (bank string)
	GetDailyRemainingAmount(c context.Context, accountNo string) (float64, error)
	ElasticSearchAccountByAccountNo(c context.Context, account model.ElasticSearchAccount) (*[]string, error)
}
