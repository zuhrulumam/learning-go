package database

import (
	"log"
)

var cs *connection

// Query doing query with existing connection
func Query(dest interface{}, query string, args ...interface{}) error {
	if cs == nil {
		log.Println("conn empty")
	}
	conn := cs.reader()
	err := conn.query(dest, query, args...)
	return err
}

// Exec doing write operation
func Exec(query string, args ...interface{}) error {
	if cs == nil {
		log.Println("conn empty")
	}

	conn := cs.writer()
	err := conn.exec(query, args...)
	return err
}
