package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"vn7n24fzkq/backend-test/config"
	"vn7n24fzkq/backend-test/models"
)

func InitDatabase() error {
	// connect database
	dsn := config.DatabaseConfig.User + ":" + config.DatabaseConfig.Password + "@tcp(" + config.DatabaseConfig.Host + ":" + config.DatabaseConfig.Port + ")/" + config.DatabaseConfig.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	print(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// NOTE: AutoMigrate will create tables, missing foreign keys, constraints, columns and indexes.
	// It will change existing column’s type if its size, precision, nullable changed.
	// It WON’T delete unused columns to protect your data.
	db.AutoMigrate(&models.Account{}, &models.User{}, &models.TodoList{})
	return nil
}
