package conf

import (
	"OpenMall/cache"
	"OpenMall/db"
)

func Init() {
	// 初始化各种配置
	conf := InitConfig()
	db.InitDB(conf)
	cache.InitRedis(conf)
}
