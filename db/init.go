package db

import (
	"OpenMall/conf"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var DB *gorm.DB

func InitDB(conf *conf.Config) *gorm.DB {
	conn, err := gorm.Open(mysql.Open(conf.SqlConn.Dsn), &gorm.Config{
		// 设置表名不会自动加复数
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("数据连接失败1！！！")
		fmt.Println(err)
		return nil
	}

	sqlDB, err := conn.DB()
	if err != nil {
		return nil
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(conf.SqlConn.Maxidleconns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(conf.SqlConn.Maxopenconns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	DB = conn
	return conn
}
