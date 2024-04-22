package dao

import (
	"OpenMall/db/model"
	"OpenMall/logger"
	"OpenMall/storage/mysql"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitMysql(cfg *mysql.Config) *gorm.DB {
	db = mysql.NewMySQL(cfg)
	AutoMigration()
	return db
}

func AutoMigration() {
	err := db.AutoMigrate(
		&model.Sku{},
		&model.Spu{},
		&model.User{},
	)
	if err != nil {
		logger.ServiceLogger.Error("AutoMigrate error", zap.Error(err))
		return
	}
}
