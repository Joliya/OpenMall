package model

import "OpenMall/db"

type Sku struct {
	db.Model
	SpuID          int    `gorm:"comment:通用信息id;not null" json:"spuID"`
	Name           string `gorm:"comment:规格名称;type:varchar(200)" json:"name"`
	Desc           string `gorm:"comment:规格详细描述;type:varchar(200)" json:"desc"`
	BasePriceFen   int    `gorm:"not null;comment:商品成本价，单位分" json:"basePrice"`
	PriceFen       int    `gorm:"not null;comment:商品售价，单位分" json:"price"`
	HeadImage      string `gorm:"not null;comment:商品头图;type:varchar(200)" json:"headImage"`
	CarouselImages string `gorm:"comment:商品轮播图列表，json字符串;type:varchar(800)" json:"carouselImages"`
	Stock          int    `gorm:"not null;comment:商品库存" json:"stock"`
}

type Spu struct {
	db.Model
	Name  string `gorm:"not null;comment:商品通用名;type:varchar(200)" json:"name"`
	Alias string `gorm:"comment:商品别名;type:varchar(200)" json:"alias"`
	Title string `gorm:"comment:商品标题，展示用;type:varchar(200)" json:"title"`
	Type  uint   `gorm:"not null;comment:商品类型" json:"type"`
}

type User struct {
	db.Model
	Username string `json:"username" gorm:"index"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
}
