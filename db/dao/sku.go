package dao

import (
	"OpenMall/db/model"
	"gorm.io/gorm"
)

type SkuDao struct {
	*gorm.DB
}

func NewSkuDao() *SkuDao {
	return &SkuDao{NewDBClient()}
}

func (dao *SkuDao) GetSkuDetail(skuID int) (r *model.Sku, err error) {
	err = dao.DB.Debug().Model(&model.Sku{}).Where("id = ?", skuID).First(&r).Error
	return
}
