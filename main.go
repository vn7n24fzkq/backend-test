package main

import (
	"vn7n24fzkq/backend-test/config"
	"vn7n24fzkq/backend-test/dao"
	"vn7n24fzkq/backend-test/database"
	"vn7n24fzkq/backend-test/routes"
)

func main() {
	// Initialize config
	config.InitConfig()
	// Initialize database
	db, err := database.InitDatabase()
	if err != nil {
		panic("failed to connect database")
	}
	dao.SetDatasource(db)
	// Initialize router
	routes.InitRouter()
}
