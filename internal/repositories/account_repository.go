package repository

import (
	"context"
	"database/sql"

	model "github.com/gamepkw/accounts-banking-microservice/internal/models"

	"github.com/go-redis/redis"
)

type accountRepository struct {
	conn  *sql.DB
	redis *redis.Client
}

func NewaccountRepository(conn *sql.DB, redis *redis.Client) AccountRepository {
	return &accountRepository{
		conn:  conn,
		redis: redis,
	}
}

type AccountRepository interface {
	GetAllAccount(ctx context.Context, cursor string, num int64) (res []model.Account, nextCursor string, err error)
	GetAccountFromRedisByAccountNo(ctx context.Context, account_no string) (*model.Account, error)
	GetAccountByAccountNo(ctx context.Context, account_no string) (*model.Account, error)
	UpdateAccount(ctx context.Context, ar *model.Account) error
	RegisterAccount(ctx context.Context, a *model.Account) error
	GetCountAccountByStatus(ctx context.Context) (result map[string]int, err error)
	DeleteAccount(ctx context.Context, account_no string) error
	GetAllAccountByUuid(ctx context.Context, uuid string) (res *[]model.Account, err error)
}
