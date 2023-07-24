package sku_service

import (
	"OpenMall/db/sku"
	"encoding/json"
)

type specification struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type GoodsDetail struct {
	Id             uint            `json:"id"`
	Name           string          `json:"name"`
	Alias          string          `json:"alias"`
	Type           uint            `json:"type"`
	BasePriceFen   int             `json:"basePrice"`
	PriceFen       int             `json:"price"`
	HeadImage      string          `json:"headImage"`
	CarouselImages []string        `json:"carouselImages"`
	Specification  []specification `json:"specification"`
	Stock          int             `json:"stock"`
}

// GetSkuDetail
// @Description: 获取商品详情
// @param skuId 商品id
// @return nil or GoodsDetail
func GetSkuDetail(skuId int) any {
	detail, err := sku.GetSkuDetail(skuId)
	if err != nil {
		return nil
	}

	return GoodsDetail{
		Id:             detail.ID,
		Name:           detail.Name,
		Alias:          detail.Alias,
		Type:           detail.Type,
		BasePriceFen:   detail.BasePriceFen,
		PriceFen:       detail.PriceFen,
		HeadImage:      detail.HeadImage,
		CarouselImages: getCarouselImages(detail.CarouselImages),
		Specification:  getSpecifications(detail.Specification),
		Stock:          detail.Stock,
	}
}

// getSpecifications
// @Description: 获取商品所有规格信息
// @param specificationStr  商品图片列表字符串，是个json字符串
// @return []specification  规格详情列表
func getSpecifications(specificationStr string) []specification {
	var spes []specification
	if specificationStr != "" {
		var speIds []int
		err := json.Unmarshal([]byte(specificationStr), &speIds)
		if err != nil {
			return nil
		}
		var speList []sku.Specification
		speList = sku.GetSpecificationList(speIds)
		for _, spe := range speList {
			spes = append(spes, specification{
				Id:   spe.ID,
				Name: spe.Name,
				Desc: spe.Desc,
			})
		}
	}
	return spes
}

// getCarouselImages
// @Description: 获取轮播图，将字符串类型的 切片 转切片
// @param carouseImagesStr    商品轮播图列表字符串，是个json字符串
// @return []string
func getCarouselImages(carouseImagesStr string) []string {
	var carouselImages []string
	if carouseImagesStr != "" {
		err := json.Unmarshal([]byte(carouseImagesStr), &carouselImages)
		if err != nil {
			carouselImages = []string{}
		}
	}
	return carouselImages
}
