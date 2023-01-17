package main

import (
	"golang_backend/controllers"
	"golang_backend/initializers"
	"golang_backend/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()
	r.POST("/post", controllers.PostsCreate)
	r.GET("/post",middleware.RequireAuth,controllers.PostGetAll)
	r.GET("/post/:id",controllers.PostGet)
	r.PUT("/post/:id",controllers.PostUpdate)
	r.DELETE("/post/:id",controllers.PostDelete)
	
	r.POST("/signup",controllers.UserCreate)
	r.POST("/login",controllers.UserLogin)


	r.Run() 
}

	