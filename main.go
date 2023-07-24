package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.String(200, "欢迎来到 OpenMall")
	})

	// 初始化数据库连接和路由
	InitConfig(router)

	defer DeferClose()

	err := router.Run()
	if err != nil {
		return
	}
}
