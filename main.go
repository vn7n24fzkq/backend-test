package main

import (
	"vn7n24fzkq/backend-test/config"
	"vn7n24fzkq/backend-test/database"
	"vn7n24fzkq/backend-test/routes"
)

func main() {
	// Initialize config
	config.InitConfig()
	// Initialize database
	database.InitDatabase()
	// Initialize router
	routes.InitRouter()
}
