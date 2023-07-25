package sku

import (
	"github.com/gin-gonic/gin"
)

func InitSkuRouters(e *gin.Engine) {
	group := e.Group("/sku")
	{
		group.GET("/detail", GetSkuDetailView)
	}
}
