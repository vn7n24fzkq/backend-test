package main

import (
	"os"
	"vn7n24fzkq/backend-test/database"
	"vn7n24fzkq/backend-test/server"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("dsn")
	var dbConn gorm.Dialector

	if dsn == "" {
		dbConn = sqlite.Open("../sqlitdb.db")
	} else {
		dbConn = mysql.Open(dsn)
	}
	// Initialize database
	db, err := database.InitDatabase(dbConn)
	if err != nil {
		panic("failed to connect database")
	}
	database.Migrate(db)

	// Create http server
	gin := gin.Default()

	server := server.InstanceServer(gin, db)
	server.Run(":8080")
}
