package sku

import "OpenMall/db"

func GetSkuDetail(skuId int) (Sku, error) {
	var sku Sku
	err := db.DB.Debug().First(&sku, skuId).Error
	return sku, err
}

func GetSpecificationList(speIds []int) []Specification {
	var speList []Specification
	db.DB.Model(&speList).Scopes(db.ArrayFilter("id", speIds, "in")).Find(&speList)
	return speList
}
