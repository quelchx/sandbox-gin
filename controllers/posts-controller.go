package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/quelchx/go-crud/initializers"
	model "github.com/quelchx/go-crud/models"
)

func PostsCreate(c *gin.Context) {
	// get the post data from the request body
	var post model.Post

	// validate the post data
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// save the post data in the database
	initializers.DB.Create(&post)

	// return a response
	c.JSON(200, gin.H{"data": post})
}

func GetPostById(c *gin.Context) {
	var post model.Post

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(400, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(200, gin.H{"data": post})
}

func GetPosts(c *gin.Context) {
	var posts []model.Post

	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{"data": posts})
}

func UpdatePostById(c *gin.Context) {
	var post model.Post

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(400, gin.H{"error": "Record not found!"})
		return
	}

	// get the post data from the request body
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Save(&post)

	c.JSON(200, gin.H{"data": post})
}

func DeletePostById(c *gin.Context) {
	var post model.Post

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(400, gin.H{"error": "Record not found!"})
		return
	}

	initializers.DB.Delete(&post)

	c.JSON(200, gin.H{"data": true})
}
