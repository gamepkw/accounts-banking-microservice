package repository

import (
	"context"
	"fmt"
	"time"

	model "github.com/gamepkw/accounts-banking-microservice/internal/models"
)

func (m *accountRepository) UpdateAccount(ctx context.Context, ar *model.Account) (err error) {
	query := `UPDATE banking.accounts set balance=?, updated_at=? WHERE account_no = ?`

	stmt, err := m.conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	*ar.UpdatedAt = time.Now()

	res, err := stmt.ExecContext(ctx, ar.Balance, ar.UpdatedAt, ar.AccountNo)
	if err != nil {
		return
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return
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
			return err
		}
		return err
	}

	return
}
