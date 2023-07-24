package main

import (
	"OpenMall/db/sku"
	"fmt"
	"gorm.io/gorm"
)

func InitMySqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		sku.Sku{},
		sku.Specification{},
	)
	if err != nil {
		fmt.Printf("AutoMigrate 失败：%s\n", err)
	}
}
