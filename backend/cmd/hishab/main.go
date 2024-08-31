package main

import (
	"github.com/gin-gonic/gin"
	"hishab.com/main/config"
	"hishab.com/main/controllers"
)

func main() {

	config.ConnectDatabase()

	r := gin.Default()

	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", controllers.Register)
		userGroup.POST("/login", controllers.Login)
	}

	categoryGroup := r.Group("/category")
	{
		categoryGroup.GET("/", controllers.FetchAllCategory)
		categoryGroup.POST("/", controllers.AddCategory)
	}

	r.Run(":8080")
}
