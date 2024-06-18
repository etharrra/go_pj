package controllers

import (
	"github.com/etharrra/go-gin/initializer"
	"github.com/etharrra/go-gin/models"
	"github.com/gin-gonic/gin"
)

func PostGet(c *gin.Context) {
	// Get all records
	var posts []models.Post
	result := initializer.DB.Find(&posts)

	if result.Error != nil {
		c.Status(500)
		return
	}

	// return the data
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostGetById(c *gin.Context) {
	// Get Id from URL
	id := c.Param("id")

	// Get all records
	var post models.Post
	result := initializer.DB.First(&post, id)

	if result.Error != nil {
		c.Status(404)
		return
	}

	// return the data
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostCreate(c *gin.Context) {
	// Get data from request
	var req struct {
		Title string `form:"title"`
		Body  string `form:"body"`
	}
	c.Bind(&req)

	// Create new record
	post := models.Post{Title: req.Title, Body: req.Body}
	result := initializer.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return the data
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// Get the id from URL
	id := c.Param("id")

	// Get data from Request
	var req struct {
		Title string `form:"title"`
		Body  string `form:"body"`
	}
	c.Bind(&req)

	// Find & Update
	var post models.Post
	initializer.DB.First(&post, id)
	initializer.DB.Model(&post).Updates(models.Post{
		Title: req.Title,
		Body:  req.Body,
	})

	// Return the data
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	// Get the id from URL
	id := c.Param("id")

	// Find & Delete
	initializer.DB.Delete(&models.Post{}, id)

	// Return the data
	c.Status(200)
}
