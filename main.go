package main

import (
	"golang_backend/controllers"
	"golang_backend/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()
	r.POST("/post", controllers.PostsCreate)
	r.GET("/post",controllers.PostGetAll)
	r.GET("/post/:id",controllers.PostGet)
	r.PUT("/post/:id",controllers.PostUpdate)
	r.DELETE("/post/:id",controllers.PostDelete)



	r.Run() 
}