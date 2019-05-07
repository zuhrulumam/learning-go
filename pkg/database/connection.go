package database

import (
	_ "github.com/go-sql-driver/mysql" // for driver
	"github.com/jinzhu/gorm"
)

type sqlQ interface {
	query(dest interface{}, query string, args ...interface{}) error
	exec(query string, args ...interface{}) error
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
		s: &mysqlgorm{
			db: db,
		},
		quitFn: closeFn,
	}

	return mysqlconn, nil

}
