package database

import (
	"database/sql"
	"fmt"
)

type mysql struct {
	db *sql.DB
}

func (m *mysql) query(q string) (*sql.Rows, error) {

	fmt.Println("query from mysql", q)

	return nil, nil
}
