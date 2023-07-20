package db

import (
	"OpenMall/conf"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB(conf *conf.Config) {
	db, err := gorm.Open("mysql", conf.SqlConn)
	db.LogMode(true)
	// Error
	if err != nil {
		fmt.Println(err)
	}
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	//默认不加复数
	db.SingularTable(true)
	//设置连接池
	////空闲
	//db.DB().SetMaxIdleConns(20)
	////打开
	//db.DB().SetMaxOpenConns(100)
	////超时
	//db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db
}
