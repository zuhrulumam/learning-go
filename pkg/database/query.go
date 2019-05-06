package database

import (
	"database/sql"
	"log"
)

var cs *connection

// QueryGorm doing query with existing connection
func QueryGorm(dest interface{}, query string, args ...interface{}) error {
	if cs == nil {
		log.Println("conn empty")
	}
	conn := cs.readerGorm()
	conn.query(dest, query, args...)
	return nil
}

// ExecGorm doing write operation
func ExecGorm(dest interface{}, query string, args ...interface{}) error {
	if cs == nil {
		log.Println("conn empty")
	}

	conn := cs.writerGorm()
	conn.exec(dest, query, args...)
	return nil
}

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
