package repository

import "context"

type SQL interface {
	ExecuteContext(context.Context, string, ...interface{}) error
	QueryContext(context.Context, string, ...interface{}) error
	QueryRowContext(context.Context, string, ...interface{}) error
	BeginTx(context.Context, string, ...interface{}) error
}

type Rows interface {
	Scan(dest ...interface{}) error
	Next() bool
	Err() error
	Close() error
}

type Row interface {
	Scan(dest ...interface{}) error
}

type Tx interface {
	ExecuteContext(context.Context, string, ...interface{}) error
	QueryContext(context.Context, string, ...interface{}) (Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) Row
	Commit() error
	Rollback() error
}
