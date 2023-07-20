package db

import (
	"OpenMall/conf"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)

func InitDB(conf *conf.Config) *gorm.DB {
	db, err := gorm.Open("mysql", conf.SqlConn)
	db.LogMode(true)
	// Error
	if err != nil {
		panic(err)
	}
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	//默认不加复数
	db.SingularTable(true)
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(20)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	return db
}
