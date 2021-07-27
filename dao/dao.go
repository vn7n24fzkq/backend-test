package dao

import "gorm.io/gorm"

var db *gorm.DB

func SetDatasource(d *gorm.DB) {
	db = d
}
