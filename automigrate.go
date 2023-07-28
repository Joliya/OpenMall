package main

import (
	"OpenMall/db/model"
	"OpenMall/util/string_util"
	"fmt"
	"gorm.io/gorm"
)

func InitMySqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.Sku{},
		model.Spu{},
		model.User{},
	)
	if string_util.IsNotNil(err) {
		fmt.Printf("AutoMigrate 失败：%s\n", err)
	}
}
