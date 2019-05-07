package database

import (
	"log"

	"github.com/jinzhu/gorm"
)

type mysqlgorm struct {
	db *gorm.DB
}

func (m *mysqlgorm) query(dest interface{}, q string, args ...interface{}) error {
	log.Println(dest)
	if dest != nil {
		m.db.Raw(q, args...).Scan(dest)
	} else {
		m.db.Raw(q, args...)
	}

	return nil
}

func (m *mysqlgorm) exec(q string, args ...interface{}) error {
	m.db.Exec(q, args...)

	return nil
}
