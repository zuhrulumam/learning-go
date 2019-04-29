package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // for driver
)

type sqlQ interface {
	query(query string) (*sql.Rows, error)
}

type connection struct {
	s sqlQ
}

func (c *connection) reader() sqlQ {
	return c.s
}

func newConnection(dbURL string) (*connection, error) {
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	mysqlconn := &connection{
		s: &mysql{
			db: db,
		},
	}

	pgconn := &connection{
		s: &pgSQL{
			db: db,
		},
	}

	fmt.Println(mysqlconn.s.query("test"))
	fmt.Println(pgconn.s.query("test"))

	return mysqlconn, nil

}
