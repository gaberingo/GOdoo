package base

import (
	"gorm.io/gorm"
)

type Model interface {
	GetID() uint
	SetID(id uint)
}

type BaseModels struct {
	gorm.Model
}

func (m *BaseModels) GetID() uint {
	return m.ID
}

func (m *BaseModels) SetID(id uint) {
	m.ID = id
}
