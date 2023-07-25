package main

import (
	"OpenMall/db/sku"
	"OpenMall/db/user"
	"OpenMall/util/string_util"
	"fmt"
	"gorm.io/gorm"
)

func InitMySqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		sku.Sku{},
		sku.Specification{},
		user.User{},
	)
	if !string_util.IsNil(err) {
		fmt.Printf("AutoMigrate 失败：%s\n", err)
	}
}
