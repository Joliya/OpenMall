/**
 * @Author: jinpeng zhang
 * @Date: 2023/7/28 14:50
 * @Description:
 */

package dao

import (
	"OpenMall/db/model"
)

type SpuDao struct {
}

func (spuDao *SpuDao) GetSpuById(spuID int) (spu *model.Spu, err error) {
	err = db.Where("id = ?", spuID).First(&spu).Error
	return
}
