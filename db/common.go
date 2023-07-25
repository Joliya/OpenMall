package db

import (
	"gorm.io/gorm"
	"reflect"
)

func ArrayFilter(key string, value interface{}, operator string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		returnDb := db
		if reflect.ValueOf(value).Len() != 0 {
			whereMap := map[string]interface{}{
				key: value,
			}
			switch operator {
			case "in":
				returnDb = returnDb.Where(whereMap)
			case "not in":
				returnDb = returnDb.Not(whereMap)
			}
		}
		return returnDb
	}
}
