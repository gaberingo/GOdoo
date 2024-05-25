package base

import (
	"gorm.io/gorm"
)

func CreateTable(db *gorm.DB, model interface{}) {
	db.AutoMigrate(model)
}
