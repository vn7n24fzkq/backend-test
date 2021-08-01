package database

import (
	"os"
	"testing"

	"gorm.io/driver/sqlite"
)

func TestInitAndMigrateDatabase(t *testing.T) {
	// We can use a random name to make all tests can be parallelizable
	tempDB := t.TempDir() + "/test.db"
	os.Remove(tempDB) // Make sure remove old test-database before test start

	// Get test database connection
	db, err := InitDatabase(sqlite.Open(tempDB))
	if err != nil {
		t.Fatal("Should not get error when init database")
	}

	migrateErr := Migrate(db)
	if migrateErr != nil {
		t.Fatal("Should not get error when migeate database")
	}
}
