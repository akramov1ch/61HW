package main

import (
	"61HW/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Tasks API
// @version 1.0
// @description This is a sample tasks API.

// @host localhost:8080
// @BasePath /

// @schemes http
func main() {
	r := routes.SetupRouter()

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
