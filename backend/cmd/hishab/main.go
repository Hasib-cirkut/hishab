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
	}

	r.Run(":8080")
}
