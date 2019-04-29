package database

import (
	"database/sql"
	"fmt"
)

type pgSQL struct {
	db *sql.DB
}

func (m *pgSQL) query(q string) (*sql.Rows, error) {

	fmt.Println("query from pgsql")

	return nil, nil
}
