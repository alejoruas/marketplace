package database

import "database/sql"

type postgresRow struct {
	row *sql.Row
}

func newPostgresRow(row *sql.Row) postgresRow {
	return postgresRow{row: row}
}

func (pr postgresRow) Scan(dest ...interface{}) error {
	if err := pr.row.Scan(dest...); err != nil {
		return err
	}

	return nil
}

type postgresRows struct {
	rows *sql.Rows
}

func newPostgresRows(rows *sql.Rows) postgresRows {
	return postgresRows{rows: rows}
}

func (pr postgresRows) Scan(dest ...interface{}) error {
	if err := pr.rows.Scan(dest...); err != nil {
		return err
	}

	return nil
}

func (pr postgresRows) Next() bool {
	return pr.rows.Next()
}

func (pr postgresRows) Err() error {
	return pr.rows.Err()
}

func (pr postgresRows) Close() error {
	return pr.rows.Close()
}
