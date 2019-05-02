package database

import (
	"database/sql"
	"log"
)

var cs *connection

// Query doing query with existing connection
func Query(query string) (*sql.Rows, error) {
	if cs == nil {
		log.Println("conn empty")
	}
	conn := cs.reader()
	return conn.query(query)
}

// Exec doing write operation
func Exec(query string) (sql.Result, error) {
	if cs == nil {
		log.Println("conn empty")
	}

	conn := cs.writer()
	return conn.exec(query)
}
