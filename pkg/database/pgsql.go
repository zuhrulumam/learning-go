package database

import (
	"database/sql"
	"fmt"
)

type pgSQL struct {
	db *sql.DB
}

func (m *pgSQL) query(q string, args ...interface{}) (*sql.Rows, error) {

	fmt.Println("query from pgsql")

	return nil, nil
}

func (m *pgSQL) exec(q string, args ...interface{}) (sql.Result, error) {

	fmt.Println("query from pgsql")

	return nil, nil
}
