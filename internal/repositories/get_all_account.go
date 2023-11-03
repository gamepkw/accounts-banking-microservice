package repository

import (
	"context"

	model "github.com/gamepkw/accounts-banking-microservice/internal/models"
	repository "github.com/gamepkw/accounts-banking-microservice/internal/repositories/helper"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (m *accountRepository) getAllAccount(ctx context.Context, query string, args ...interface{}) (accounts []model.Account, err error) {

	rows, err := m.conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	accounts = make([]model.Account, 0)

	for rows.Next() {
		account := model.Account{}

		err = rows.Scan(
			&account.AccountNo,
			&account.Uuid,
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

func (m *accountRepository) GetAllAccount(ctx context.Context, cursor string, num int64) (res []model.Account, nextCursor string, err error) {
	query := `SELECT * FROM banking.accounts WHERE created_at > ? ORDER BY created_at LIMIT ? `

	decodedCursor, err := repository.DecodeCursor(cursor)
	if err != nil && cursor != "" {
		return nil, "", model.ErrBadParamInput
	}

	res, err = m.getAllAccount(ctx, query, decodedCursor, num)
	if err != nil {
		return nil, "", err
	}

	if len(res) == int(num) {
		nextCursor = repository.EncodeCursor(*res[len(res)-1].CreatedAt)
	}

	return
}
