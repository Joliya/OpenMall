package main

import (
	"OpenMall/conf"
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

	conf.Init()
}
