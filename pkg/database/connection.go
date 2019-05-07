package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // for driver
)

type sqlQ interface {
	query(query string, args ...interface{}) (*sql.Rows, error)
	exec(query string, args ...interface{}) (sql.Result, error)
}

type connection struct {
	s      sqlQ
	quitFn func()
}

func (c *connection) reader() sqlQ {
	return c.s
}

func (c *connection) writer() sqlQ {
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

	closeFn := func() {
		db.Close()
	}

	mysqlconn := &connection{
		s: &mysql{
			db: db,
		},
		quitFn: closeFn,
	}

	pgconn := &connection{
		s: &pgSQL{
			db: db,
		},
		quitFn: closeFn,
	}

	fmt.Println(pgconn.s.query("test"))

	return mysqlconn, nil

}
