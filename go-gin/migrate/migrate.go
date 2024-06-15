package main

import (
	"github.com/etharrra/go-gin/initializer"
	"github.com/etharrra/go-gin/models"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectDB()
}

func main() {
	// Migrate the schema
	initializer.DB.AutoMigrate(&models.Post{})
}
