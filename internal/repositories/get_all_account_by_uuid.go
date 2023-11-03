package repository

import (
	"context"

	model "github.com/gamepkw/accounts-banking-microservice/internal/models"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (m *accountRepository) GetAllAccountByUuid(ctx context.Context, uuid string) (res *[]model.Account, err error) {
	accounts, err := m.fetchAllAccountFromDatabaseByUuid(ctx, uuid)
	if err != nil {
		return nil, err
	}
	return &accounts, nil
}

func (m *accountRepository) fetchAllAccountFromDatabaseByUuid(ctx context.Context, uuid string) (accounts []model.Account, err error) {
	query := `SELECT * FROM banking.accounts WHERE uuid = ?`

	rows, err := m.conn.QueryContext(ctx, query, uuid)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	accounts = make([]model.Account, 0)

	for rows.Next() {
		account := model.Account{}

		err = rows.Scan(
			&account.AccountNo,
			&account.Uuid,
			&account.Name,
			&account.Email,
			&account.Tel,
			&account.Balance,
			&account.Bank,
			&account.Status,
			&account.IsClosed,
			&account.CreatedAt,
			&account.UpdatedAt,
		)
		if err != nil {
			logrus.Error(err)
			return accounts, errors.Wrap(err, "error scan")
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}
