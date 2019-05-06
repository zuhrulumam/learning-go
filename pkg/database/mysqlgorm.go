package database

import (
	"github.com/jinzhu/gorm"
)

type mysqlgorm struct {
	db *gorm.DB
}

func (m *mysqlgorm) query(dest interface{}, q string, args ...interface{}) error {

	m.db.Raw(q, args...).Scan(dest)
	return nil
}

func (m *mysqlgorm) exec(dest interface{}, q string, args ...interface{}) error {
	m.db.Exec(q, args...).Scan(dest)
	return nil
}
