package main

import (
	"github.com/gin-gonic/gin"

	"inheritence_management/controllers"
	"inheritence_management/middlewares"
	"inheritence_management/models"

	docs "inheritence_management/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Inheritence Management
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// host localhost:8080
// @host 54.238.205.73
// @BasePath /api

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	models.ConnectDataBase()
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	auth := r.Group("/api/auth")
	auth.Use(middlewares.CORSMiddleware())
	auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)

	api := r.Group("/api")
	api.Use(middlewares.JwtAuthMiddleware())
	api.GET("/profile", controllers.CurrentUser)
	api.GET("/user", controllers.GetAllUser)
	api.GET("/family", controllers.GetRelationship)
	api.POST("/family", controllers.AddRelationship)
	api.GET("/image", controllers.GetImage)
	api.POST("/image", controllers.UploadImage)
	api.GET("/package", controllers.GetPackage)
	api.POST("/package", controllers.CreatePackage)

	// sess := storage.ConnectAws()
	// image := r.Group("/api/image")
	// image.Use(func(c *gin.Context) {
	// 	c.Set("sess", sess)
	// 	c.Next()
	// })

	r.Run(":8080")
}
