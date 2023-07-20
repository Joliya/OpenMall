package main

import (
	"OpenMall/conf"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.String(200, "欢迎来到 OpenMall")
	})
	err := router.Run()
	if err != nil {
		return
	}

	//初始化配置
	config := conf.InitConfig()
	//初始化mysql
	sqlConn := config.SqlConn
	fmt.Println(sqlConn)

	//初始化redis
	fmt.Println("redis db num:", config.RedisConf.Db)

	//初始化路由
	route := gin.Default()

	//项目启动
	_ = route.Run(config.AppConf.Port)
}
