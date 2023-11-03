package repository

import (
	"context"

	model "github.com/gamepkw/accounts-banking-microservice/internal/models"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (m *accountRepository) GetCountAccountByStatus(ctx context.Context) (result map[string]int, err error) {
	query := `SELECT status_list.status, COALESCE(status_count.count, 0) AS count
	FROM (
		SELECT 'active' AS status
		UNION SELECT 'inactive'
		UNION SELECT 'fraud'
		UNION SELECT 'zero'
	) AS status_list
	LEFT JOIN (
		SELECT status, COUNT(*) AS count
		FROM banking.accounts
		WHERE status IN ('active', 'inactive', 'fraud', 'zero')
		GROUP BY status
	) AS status_count ON status_list.status = status_count.status`
	rows, err := m.conn.QueryContext(ctx, query)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	statusCounts := make(map[string]int)

	for rows.Next() {
		countAccount := model.CountAccount{}

		err = rows.Scan(
			&countAccount.Status,
			&countAccount.Count,
		)

		if err != nil {
			return nil, errors.Wrap(err, "error scan")
		}

		statusCounts[countAccount.Status] = countAccount.Count
	}

	return statusCounts, nil
}
