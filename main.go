package main

import (
	"fmt"

	"GOdoo.gabe/apotek"
	"GOdoo.gabe/base"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("db_godoo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// db.Create(apotek.CreateObat())
	var result []apotek.Obat

	err = base.SearchRecord(db, "", &result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	for _, res := range result {
		fmt.Println(res.Name)
	}
	// for _, res := range result {
	// 	fmt.Println(res)
	// }
	// base.CreateTable(db, &apotek.Obat{})

	// db.Find(&obat)
	// for _, ob := range obat {
	// 	fmt.Println(ob.ID)
	// }
}
