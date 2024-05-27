package base

import (
	"gorm.io/gorm"
)

func CreateTable(db *gorm.DB, model interface{}) {
	db.AutoMigrate(model)
}

func SearchRecord(db *gorm.DB, query string, table interface{}) error {
	err := db.Where(query).Find(table).Error
	if err != nil {
		return err
	}
	return nil
}
