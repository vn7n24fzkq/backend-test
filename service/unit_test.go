package service

import (
	"os"
	"testing"
	"vn7n24fzkq/backend-test/database"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// We write some of common function in this file for tests under package

func GetTestDB(t *testing.T) *gorm.DB {
	// We can use a random name to make all tests can be parallelizable
	tempDB := t.TempDir() + "/test.db"
	os.Remove(tempDB) // Make sure remove old test-database before test start

	// Get test database connection
	db, err := database.InitDatabase(sqlite.Open(tempDB))
	if err != nil {
		panic("failed to connect test database")
	}
	database.Migrate(db)
	return db
}
