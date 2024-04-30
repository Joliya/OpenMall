package sku

import (
	"OpenMall/http_tools"
	"OpenMall/util/string_util"
	"github.com/gin-gonic/gin"
	"strconv"
)
import "OpenMall/services/sku_service"

type SkuDto struct {
	SkuId             int
	SpuId             int
	Name              string
	Alias             string
	Title             string
	Type              int
	SpecificationName string
	Desc              string
	BasePrice         string
	Price             string
	HeadImage         string
	CarouselImages    []string
	Stock             int
}

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

func CreateSkuView(c *gin.Context) {
	// 将 c 中的数据解析为 SkuDto
	sku := SkuDto{}
	err := c.BindJSON(&sku)
	if err != nil {
		http_tools.JsonHttpError(c, "参数错误", nil)
	}
	// 调用 sku_service.CreateSku
	// 返回结果
}

func UpdateSkuView(c *gin.Context) {
	// 将 c 中的数据解析为 SkuDto
	sku := SkuDto{}
	err := c.BindJSON(&sku)
	if err != nil {
		http_tools.JsonHttpError(c, "参数错误", nil)
	}
	// 调用 sku_service.UpdateSku
	// 返回结果
}

func DeleteSkuView(c *gin.Context) {
	skuId, exist := c.GetQuery("skuId")
	if !exist {
		http_tools.JsonHttpError(c, "商品不存在", nil)
	}
	_, err := strconv.Atoi(skuId)
	if string_util.IsNotNil(err) {
		http_tools.JsonHttpError(c, "商品不存在", nil)
	}
	// 调用 sku_service.DeleteSku
	// 返回结果
}
