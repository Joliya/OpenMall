package sku_service

import (
	"OpenMall/db/dao"
	"OpenMall/util/number"
	"OpenMall/util/string_util"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type spu struct {
	ID        uint
	Name      string
	Alias     string
	Title     string
	Type      uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type goodsDetail struct {
	ID                uint
	SpuID             int
	Name              string
	Alias             string
	Title             string
	Type              uint
	SpecificationName string
	Desc              string
	BasePrice         string
	Price             string
	HeadImage         string
	CarouselImages    []string
	Stock             int
	CreatedAt         string
	UpdatedAt         string
	DeletedAt         string
}

// GetSkuDetail
// @Description: 获取商品详情
// @param skuId 商品id
// @return nil or GoodsDetail
func GetSkuDetail(skuId int) any {
	skuDao := dao.NewSkuDao()
	detail, err := skuDao.GetSkuDetail(skuId)
	if string_util.IsNotNil(err) {
		return nil
	}
	spuDao := dao.NewSpuDao()
	spu, err := spuDao.GetSpuById(detail.SpuID)
	return goodsDetail{
		ID:                detail.ID,
		SpuID:             detail.SpuID,
		Name:              spu.Name,
		Alias:             spu.Alias,
		Title:             spu.Title,
		Type:              spu.Type,
		SpecificationName: detail.Name,
		Desc:              detail.Desc,
		BasePrice:         number.GetYuanFromFen(detail.BasePriceFen),
		Price:             number.GetYuanFromFen(detail.PriceFen),
		HeadImage:         detail.HeadImage,
		CarouselImages:    getCarouselImages(detail.CarouselImages),
		Stock:             detail.Stock,
		CreatedAt:         detail.CreatedAt.Format(time.DateTime),
		UpdatedAt:         detail.UpdatedAt.Format(time.DateTime),
		DeletedAt:         detail.DeletedAt.Time.Format(time.DateTime),
	}
}

// getCarouselImages
// @Description: 获取轮播图，将字符串类型的 切片 转切片
// @param carouseImagesStr    商品轮播图列表字符串，是个json字符串
// @return []string
func getCarouselImages(carouseImagesStr string) []string {
	var carouselImages []string
	if string_util.IsEmpty(carouseImagesStr) {
		err := json.Unmarshal([]byte(carouseImagesStr), &carouselImages)
		if string_util.IsNotNil(err) {
			carouselImages = []string{}
		}
	}
	return carouselImages
}
