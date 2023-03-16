package main

import (
	"github.com/MuShaf-NMS/SV-BE-Test/config"
	"github.com/MuShaf-NMS/SV-BE-Test/database"
	"github.com/MuShaf-NMS/SV-BE-Test/router"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config := config.GetConfig()
	db := database.CreateConnection(config)
	defer database.CloseConnection(db)
	if config.App_Mode == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()
	router.InitializeRoute(server, db, *config)
	server.Run(":" + config.App_Port)
}
