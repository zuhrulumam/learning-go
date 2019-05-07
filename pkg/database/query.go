package database

import (
	"database/sql"
	"log"
)

var cs *connection

// Query doing query with existing connection
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	if cs == nil {
		log.Println("conn empty")
	}
	conn := cs.reader()
	return conn.query(query, args...)
}

// Exec doing write operation
func Exec(query string, args ...interface{}) (sql.Result, error) {
	if cs == nil {
		log.Println("conn empty")
	}

	conn := cs.writer()
	return conn.exec(query, args...)
}
