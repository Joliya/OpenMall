package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	// 初始化数据库连接和路由
	InitConfig(r)

	err := r.Run()
	if err != nil {
		return
	}
}
