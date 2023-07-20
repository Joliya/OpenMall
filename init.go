package main

import (
	"OpenMall/cache"
	"OpenMall/conf"
	"OpenMall/db"
)

func Init() {
	// 初始化各种配置
	initConfig := conf.InitConfig()
	db.InitDB(initConfig)
	cache.InitRedis(initConfig)
}
