package main

import (
	"github.com/gin-gonic/gin"

	"inheritence_management/controllers"
	"inheritence_management/middlewares"
	"inheritence_management/models"

	//docs "github.com/inheritence_management_backend/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin Swagger Example API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	models.ConnectDataBase()
	r := gin.Default()

	url := ginSwagger.URL("http://localhost:3000/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	auth := r.Group("/api/auth")
	auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)

	api := r.Group("/api")
	api.Use(middlewares.JwtAuthMiddleware())
	api.GET("/user", controllers.CurrentUser)

	r.Run(":8080")
}
