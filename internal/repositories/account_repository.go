package repository

import (
	"context"
	"database/sql"

	model "github.com/gamepkw/accounts-banking-microservice/internal/models"

	"github.com/elastic/go-elasticsearch/v8"

	"github.com/go-redis/redis"
)

type accountRepository struct {
	conn    *sql.DB
	redis   *redis.Client
	elastic *elasticsearch.Client
}

func NewAccountRepository(conn *sql.DB, redis *redis.Client, elastic *elasticsearch.Client) AccountRepository {
	return &accountRepository{
		conn:    conn,
		redis:   redis,
		elastic: elastic,
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
	GetConfig(ctx context.Context, configName string) (string, error)
	ElasticSearchAccountByAccountNo(ctx context.Context, account model.ElasticSearchAccount) (*[]string, error)
}
