package database

import (
	"database/sql"
)

type mysql struct {
	db *sql.DB
}

func (m *mysql) query(q string) (*sql.Rows, error) {

	return m.db.Query(q)
}

func (m *mysql) exec(q string) (sql.Result, error) {
	return m.db.Exec(q)
}
