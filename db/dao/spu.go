/**
 * @Author: jinpeng zhang
 * @Date: 2023/7/28 14:50
 * @Description:
 */

package dao

import (
	"OpenMall/db/model"
	"gorm.io/gorm"
)

type SpuDao struct {
	*gorm.DB
}

func NewSpuDao() *SpuDao {
	return &SpuDao{NewDBClient()}
}

func (spuDao *SpuDao) GetSpuById(spuID int) (r *model.Spu, err error) {
	err = spuDao.DB.Model(&model.Spu{}).Where("id = ?", spuID).First(&r).Error
	return
}
