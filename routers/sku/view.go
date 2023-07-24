package sku

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
import "OpenMall/services/sku_service"

func GetSkuDetailView(c *gin.Context) {
	skuId, exist := c.GetQuery("skuId")
	if !exist {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "商品不存在", "data": nil})
	}
	skuID, err := strconv.Atoi(skuId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "商品不存在", "data": nil})
	}
	detail := sku_service.GetSkuDetail(skuID)
	if detail == nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "商品不存在", "data": nil})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "", "data": detail})
	}
}
