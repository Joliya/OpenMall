package main

import (
	"OpenMall/db/sku"
	"OpenMall/db/user"
	"fmt"
	"gorm.io/gorm"
)

func InitMySqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		sku.Sku{},
		sku.Specification{},
		user.User{},
	)
	if err != nil {
		fmt.Printf("AutoMigrate 失败：%s\n", err)
	}
}
