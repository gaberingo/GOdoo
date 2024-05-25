package apotek

import (
	"time"

	"GOdoo.gabe/base"
)

type Obat struct {
	base.Model
	Name    string
	Price   int
	Stock   int
	Expired *time.Time
}

func (o *Obat) TableName() string {
	return "obat"
}

func CreateObat() *Obat {
	return &Obat{Name: "Sanmol", Price: 10000, Stock: 10}
}
