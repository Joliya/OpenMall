package dao

import (
	"OpenMall/db/model"
)

type SkuDao struct {
}

func (dao *SkuDao) GetSkuDetail(skuID int) (sku *model.Sku, err error) {
	err = db.Where("id = ?", skuID).First(&sku).Error
	return
}
