package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

func init() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		DB = db
		log.Print("DB Connection Success")
	}
	if !DB.Migrator().HasTable(&User{}) {
		db.Migrator().CreateTable(&User{})
	}
}

func main() {
	r := gin.Default()
	r.GET("/user", GetAllUsers)
	r.GET("/user/:id", GetUserById)
	r.POST("/user", CreateUser)
	r.PUT("/user/:id", UpdateUser)
	r.DELETE("/user/:id", DeleteUser)
	r.Run()
}

func GetAllUsers(c *gin.Context) {
	var users []User
	DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"user": users,
	})
}

func GetUserById(c *gin.Context) {
	var user User
	if err := DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	DB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}

// UpdateUser handles HTTP PUT requests to update an existing user's details in the database.
func UpdateUser(c *gin.Context) {
	// Retrieve the user by ID from the database
	var user User
	if res := DB.First(&user, c.Param("id")); res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": res.Error,
		})
		return
	}

	// Bind the incoming JSON request data to a struct
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	// Update the user record in the database with the new data
	res := DB.Model(&user).Updates(User{
		Name:  req.Name,
		Email: req.Email,
	})
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, res.Error)
		return
	}

	// Return the updated user data in the HTTP response
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// DeleteUser deletes a user from the database based on the provided user ID.
func DeleteUser(c *gin.Context) {
	res := DB.Delete(&User{}, c.Param("id"))
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted",
	})
}
