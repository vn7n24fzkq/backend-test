package main

import (
	"vn7n24fzkq/backend-test/config"
	"vn7n24fzkq/backend-test/dao"
	"vn7n24fzkq/backend-test/database"
	"vn7n24fzkq/backend-test/routes"
	"vn7n24fzkq/backend-test/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
)

func main() {
	// Initialize config
	config.InitConfig()

	// Initialize database
	dsn := config.DatabaseConfig.User + ":" + config.DatabaseConfig.Password + "@tcp(" + config.DatabaseConfig.Host + ":" + config.DatabaseConfig.Port + ")/" + config.DatabaseConfig.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	dbConn := mysql.Open(dsn)
	db, err := database.InitDatabase(dbConn)
	if err != nil {
		panic("failed to connect database")
	}
	database.Migrate(db)

	// Initialize DAO
	userDAO := dao.NewUserDAO(db)
	taskDAO := dao.NewTaskDAO(db)

	// Initialize Service
	userService := service.NewUserService(userDAO)
	taskService := service.NewTaskService(taskDAO, userService)

	// Create http server
	gin := gin.Default()

	server := routes.InstanceServer(gin, userService, taskService)
	server.Run(":8080")
}
