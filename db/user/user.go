package user

import "OpenMall/db"

type User struct {
	db.Model
	Username string `json:"username" gorm:"index"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
}
