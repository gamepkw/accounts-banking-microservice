package repository

import (
	"context"
	"time"

	model "github.com/gamepkw/accounts-banking-microservice/internal/models"
)

func (m *accountRepository) RegisterAccount(ctx context.Context, a *model.Account) (err error) {
	query := `INSERT banking.accounts SET account_no=?, uuid=?, name=? , email=? , tel=?, bank=? , created_at=? , updated_at=?`
	stmt, err := m.conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, a.AccountNo, a.Uuid, a.Name, a.Email, a.Tel, a.Bank, time.Now(), time.Now())
	if err != nil {
		return
	}
	return
}
