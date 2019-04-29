package database

import (
	// "database/sql"
	"log"
)

type masterConn struct {
	conn *connection
}

func (m *masterConn) reader() sqlQ {
	return m.conn.s
}

func setEmptyConn() *masterConn {
	return &masterConn{
		conn: nil,
	}
}

func Configure(dbURL string) {
	currConn := setEmptyConn()

	conn, err := newConnection(dbURL)
	if err != nil {
		log.Println("Cannot connect database")
	}

	currConn.conn = conn

}
