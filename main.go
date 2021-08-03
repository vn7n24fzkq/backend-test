package main

import (
	"fmt"
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
		fmt.Println("-------------")
		fmt.Println("We are using Sqlite as a database")
		fmt.Println("-------------")

		dbConn = sqlite.Open("./sqliteDB.db")
	} else {
		fmt.Println("-------------")
		fmt.Println("We are using Mariadb as a database")
		fmt.Println("-------------")

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
