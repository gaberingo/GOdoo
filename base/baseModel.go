package base

import (
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	TableName   string
	Description string
}

func (m *Model) GetTableName() string {
	return m.TableName
}
