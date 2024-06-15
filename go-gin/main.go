package main

import (
	"github.com/etharrra/go-gin/controllers"
	"github.com/etharrra/go-gin/initializer"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectDB()
}

func main() {
	r := gin.Default()
	r.GET("/post", controllers.PostGet)
	r.GET("/post/:id", controllers.PostGetById)
	r.POST("/post", controllers.PostCreate)
	r.PUT("/post/:id", controllers.PostUpdate)
	r.DELETE("/post/:id", controllers.PostDelete)
	r.Run() // listen and serve on 0.0.0.0:8080
}
