package main

import (
	"OpenMall/util/string_util"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 初始化数据库连接和路由
	InitConfig(r)
	err := r.Run()
	if !string_util.IsNil(err) {
		return
	}
}
