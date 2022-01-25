package main

import (
	"bookstore/controllers"
	"bookstore/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectDatabase()

	router.GET("/books", controllers.AllBooks)
	router.POST("/books", controllers.CreateBook)
	router.GET("/books/:id", controllers.FindBook)
	router.PUT("/books/:id", controllers.UpdateBook)

	router.Run(":8080")
}
