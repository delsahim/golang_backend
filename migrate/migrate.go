package main

import (
	"golang_backend/initializers"
	"golang_backend/models"
)


func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()

}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
	initializers.DB.AutoMigrate(&models.UserModel{})

}