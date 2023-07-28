/**
 * @Author: jinpeng zhang
 * @Date: 2023/7/28 14:48
 * @Description:
 */

package model

import "OpenMall/db"

type Spu struct {
	db.Model
	Name  string `gorm:"not null;comment:商品通用名;type:varchar(200)" json:"name"`
	Alias string `gorm:"comment:商品别名;type:varchar(200)" json:"alias"`
	Title string `gorm:"comment:商品标题，展示用;type:varchar(200)" json:"title"`
	Type  uint   `gorm:"not null;comment:商品类型" json:"type"`
}
