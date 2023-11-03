package repository

import (
	"context"
	"fmt"
	"time"

	model "github.com/gamepkw/accounts-banking-microservice/internal/models"
	"github.com/pkg/errors"
)

func (m *accountRepository) UpdateAccount(ctx context.Context, ar *model.Account) (err error) {
	query := `UPDATE banking.accounts set balance=?, updated_at=? WHERE account_no = ?`

	stmt, err := m.conn.PrepareContext(ctx, query)
	if err != nil {
		return errors.Wrap(err, "error sql statement")
	}

	*ar.UpdatedAt = time.Now()

	res, err := stmt.ExecContext(ctx, ar.Balance, ar.UpdatedAt, ar.AccountNo)
	if err != nil {
		return errors.Wrap(err, "error execute sql statement")
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "error row affected")
	}
	if affect != 1 {
		err = fmt.Errorf("weird  Behavior. Total Affected: %d", affect)
		return
	}

	cacheKey := fmt.Sprintf("account_no: %s", ar.AccountNo)

	if affect == 1 {
		err := m.redis.Del(cacheKey).Err()
		if err != nil {
			fmt.Printf("Error clearing key '%s': %v\n", cacheKey, err)
			return errors.Wrap(err, fmt.Sprintf("error clear redis key: %s", cacheKey))
		}
		return err
	}

	return
}
