package database

import (
	"github.com/jinzhu/gorm"
)

type mysqlgorm struct {
	db *gorm.DB
}

func (m *mysqlgorm) query(dest interface{}, q string, args ...interface{}) error {
	err := m.db.Raw(q, args...).Scan(dest).Error
	// err := db.Error
	// if gorm.IsRecordNotFoundError(err) {
	// 	err = errors.New("Record Not found")
	// }

	// if err != nil {
	// 	log.Println("err from mysqlgorm ", err)
	// }

	return err
}

func (m *mysqlgorm) exec(q string, args ...interface{}) error {
	db := m.db.Exec(q, args...)
	err := db.Error

	return err
}
