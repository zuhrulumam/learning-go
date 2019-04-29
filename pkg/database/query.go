package database

import (
	"database/sql"
	"log"
)

var cs *masterConn

func Query(query string) (*sql.Rows, error) {
	if cs == nil {
		log.Println("master conn empty")
	}
	conn := cs.reader()
	return conn.query(query)
}
