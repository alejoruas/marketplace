package database

import (
	"context"
	"database/sql"
	"fmt"
	"marketplace/adapter/repository"

	_ "github.com/lib/pq"
)

type postgressql struct {
	db *sql.DB
}

func NewPostgressql(config *configdb) (*postgressql, error) {
	var datasource = fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		config.host, config.port, config.database, config.user, config.password)

	connection, err := sql.Open(config.driver, datasource)
	if err != nil {
		return &postgressql{}, err
	}

	err = connection.Ping()
	if err != nil {
		fmt.Println(err)
		//return &postgressql{}, err
	}

	return &postgressql{db: connection}, nil
}

func (p postgressql) BeginTx(ctx context.Context) (repository.Tx, error) {
	tx, err := p.db.BeginTx(ctx, &sql.TxOptions{})

	if err != nil {
		return postgresTx{}, err
	}

	return newPostgresTx(tx), nil
}

func (p postgressql) ExecuteContext(ctx context.Context, query string, args ...interface{}) error {
	_, err := p.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (p postgressql) QueryContext(ctx context.Context, query string, args ...interface{}) (repository.Rows, error) {
	rows, err := p.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	row := newPostgresRows(rows)

	return row, nil
}

func (p postgressql) QueryRowContext(ctx context.Context, query string, args ...interface{}) repository.Row {
	row := p.db.QueryRowContext(ctx, query, args...)

	return newPostgresRow(row)
}
