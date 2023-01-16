package controllers

import (
	"golang_backend/initializers"
	"golang_backend/models"
	"golang_backend/schemas"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context){
	c.Bind(&schemas.RequestBody)

	post:=models.Post{Title: schemas.RequestBody.Title, Body: schemas.RequestBody.Body}
	result:= initializers.DB.Create(&post)
	if result.Error!=nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})	
}

func PostGetAll(c *gin.Context) {
	var posts []models.Post

	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})	
}

func PostGet(c *gin.Context) {
	id:=c.Param("id")

	var post models.Post

	initializers.DB.First(&post,id)

	c.JSON(200, gin.H{
		"post": post,
	})	
}

func PostUpdate(c *gin.Context) {
	id:=c.Param("id")
	c.Bind(&schemas.RequestBody)

	var post models.Post
	initializers.DB.First(&post,id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: schemas.RequestBody.Title,
		Body: schemas.RequestBody.Body,
	})
	c.JSON(200, gin.H{
		"post": post,
	})	
}

func PostDelete(c *gin.Context) {
	id:=c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)
	c.Status(200)
}
