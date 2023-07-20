package db

import (
	"gorm.io/gorm"
)

type Article struct {
	Id int64 `gorm:"primaryKey,autoIncrement"`
	// 存的是作者的用户 ID
	Author  int64  `gorm:"not null"`
	Title   string `form:"title"`
	Content string `form:"content"`
	Ctime   int64
	Utime   int64
}
