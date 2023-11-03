package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	model "github.com/gamepkw/accounts-banking-microservice/internal/models"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

func (m *accountRepository) GetAccountByAccountNo(ctx context.Context, account_no string) (res *model.Account, err error) {
	account, err := m.fetchAccountFromDatabase(ctx, account_no)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (m *accountRepository) fetchAccountFromDatabase(ctx context.Context, account_no string) (res model.Account, err error) {

	query := `SELECT * FROM banking.accounts WHERE account_no = ?`

	list, err := m.getAllAccount(ctx, query, account_no)
	if err != nil {
		return model.Account{}, errors.Wrap(err, "error fetch account from database")
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, model.ErrNotFound
	}

	return
}

func (m *accountRepository) GetAccountFromRedisByAccountNo(ctx context.Context, account_no string) (*model.Account, error) {
	// Check if the user exists in Redis cache
	cacheKey := fmt.Sprintf("account_no: %s", account_no)
	cachedAccount, err := m.redis.Get(cacheKey).Result()

	if err == redis.Nil {
		// Cache miss: key does not exist in Redis
		account, err := m.fetchAccountFromDatabase(ctx, account_no)
		if err != nil {
			return nil, err
		}

		ttl := 0.0 * time.Second

		serializedAccount := m.serializeAccount(&account)
		err = m.redis.Set(cacheKey, serializedAccount, ttl).Err()
		if err != nil {
			return nil, err
		}
		return &account, nil
	} else if err != nil {
		return nil, fmt.Errorf("error parsing account from cache: %v", err)
	} else {
		// Cache hit: key exists in Redis, use the retrieved value
		account, err := m.parseAccountFromCache(cachedAccount)
		if err != nil {
			return nil, fmt.Errorf("error parsing account from cache: %v", err)
		}
		return account, nil
	}
}

func (m *accountRepository) parseAccountFromCache(cachedContact string) (*model.Account, error) {
	var account model.Account
	err := json.Unmarshal([]byte(cachedContact), &account)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing user from cache")
	}

	return &account, nil
}

func (m *accountRepository) serializeAccount(account *model.Account) string {
	jsonData, _ := json.Marshal(account)
	return string(jsonData)
}
