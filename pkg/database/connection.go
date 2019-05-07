package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // for driver
	"github.com/jinzhu/gorm"
)

type sqlQ interface {
	query(query string, args ...interface{}) (*sql.Rows, error)
	exec(query string, args ...interface{}) (sql.Result, error)
}

type sqlQGorm interface {
	query(dest interface{}, query string, args ...interface{}) error
	exec(dest interface{}, query string, args ...interface{}) error
}

type connection struct {
	s      sqlQ
	sGorm  sqlQGorm
	quitFn func()
}

func (c *connection) readerGorm() sqlQGorm {
	return c.sGorm
}

func (c *connection) writerGorm() sqlQGorm {
	return c.sGorm
}

func (c *connection) reader() sqlQ {
	return c.s
}

func (c *connection) writer() sqlQ {
	return c.s
}

func newConnectionGorm(dbURL string) (*connection, error) {
	db, err := gorm.Open("mysql", dbURL)
	if err != nil {
		return nil, err
	}

	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}

	closeFn := func() {
		db.Close()
	}
	db.Debug()

	mysqlconn := &connection{
		sGorm: &mysqlgorm{
			db: db,
		},
		quitFn: closeFn,
	}

	return mysqlconn, nil

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
