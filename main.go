package main

import (
	"GOdoo.gabe/apotek"
	// "GOdoo.gabe/base"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("db_godoo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// base.CreateTable(db, &apotek.Obat{})
	db.Create(apotek.CreateObat())
}
