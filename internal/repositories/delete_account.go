package repository

import (
	"context"
	"fmt"
	"time"
)

func (m *accountRepository) DeleteAccount(ctx context.Context, account_no string) (err error) {
	query := `UPDATE banking.accounts set updated_at=? , is_deleted = 1 WHERE account_no = ?`

	stmt, err := m.conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, time.Now(), account_no)
	if err != nil {
		return
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return
	}

	if rowsAfected != 1 {
		err = fmt.Errorf("weird  Behavior. Total Affected: %d", rowsAfected)
		return
	}

	if rowsAfected == 1 {
		err = fmt.Errorf("Delete completed")
		return
	}

	return
}
