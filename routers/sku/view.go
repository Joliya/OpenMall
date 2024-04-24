package sku

import (
	"OpenMall/http_tools"
	"OpenMall/util/string_util"
	"github.com/gin-gonic/gin"
	"strconv"
)
import "OpenMall/services/sku_service"

func GetSkuDetailView(c *gin.Context) {
	skuId, exist := c.GetQuery("skuId")
	if !exist {
		http_tools.JsonHttpError(c, "商品不存在", nil)
	}
	skuID, err := strconv.Atoi(skuId)
	if string_util.IsNotNil(err) {
		http_tools.JsonHttpError(c, "商品不存在", nil)
	}
	detail := sku_service.GetSkuDetail(skuID)
	if string_util.IsNil(detail) {
		http_tools.JsonHttpError(c, "商品不存在", nil)
	} else {
		http_tools.JsonHttpSuccess(c, detail)
	}
}
