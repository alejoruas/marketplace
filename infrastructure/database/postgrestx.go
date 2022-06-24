package database

import (
	"context"
	"database/sql"
	"marketplace/adapter/repository"
)

type postgresTx struct {
	tx *sql.Tx
}

func newPostgresTx(tx *sql.Tx) postgresTx {
	return postgresTx{tx: tx}
}

func (p postgresTx) ExecuteContext(ctx context.Context, query string, args ...interface{}) error {
	_, err := p.tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (p postgresTx) QueryContext(ctx context.Context, query string, args ...interface{}) (repository.Rows, error) {
	rows, err := p.tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	row := newPostgresRows(rows)

	return row, nil
}

func (p postgresTx) QueryRowContext(ctx context.Context, query string, args ...interface{}) repository.Row {
	row := p.tx.QueryRowContext(ctx, query, args...)

	return newPostgresRow(row)
}

func (p postgresTx) Commit() error {
	return p.tx.Commit()
}

func (p postgresTx) Rollback() error {
	return p.tx.Rollback()
}
