package main

import (
	"OpenMall/cache"
	"OpenMall/conf"
	"OpenMall/db"
	"OpenMall/routers"
	"github.com/gin-gonic/gin"
)

func InitConfig(server *gin.Engine) {
	// 初始化各种配置
	initConfig := conf.InitConfig()
	conn := db.InitDB(initConfig)
	InitMySqlTables(conn)
	cache.InitRedis(initConfig)
	// 绑定路由
	routers.InitRouter(server)
}
