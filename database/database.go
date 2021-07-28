package database

import (
	"gorm.io/gorm"
	"vn7n24fzkq/backend-test/dao"
)

func InitDatabase(dbConnector gorm.Dialector) (*gorm.DB, error) {
	db, err := gorm.Open(dbConnector, &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	// NOTE: AutoMigrate will create tables, missing foreign keys, constraints, columns and indexes.
	// It will change existing column’s type if its size, precision, nullable changed.
	// It WON’T delete unused columns to protect your data.
	err := db.AutoMigrate(&dao.User{}, &dao.Task{})
	return err
}
