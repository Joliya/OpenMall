package db

import (
	"OpenMall/conf"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func InitDB(conf *conf.Config) *gorm.DB {
	DB, err := gorm.Open(mysql.Open(conf.SqlConn))
	if err != nil {
		fmt.Println("数据连接失败1！！！")
		fmt.Println(err)
		return nil
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return nil
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return DB

}
