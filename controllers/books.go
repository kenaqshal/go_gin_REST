package controllers

import (
	"bookstore/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AllBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

func CreateBook(c *gin.Context) {
	var inputBook models.CreateBookInput
	if err := c.ShouldBindJSON(&inputBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{
		Title:  inputBook.Title,
		Author: inputBook.Author,
	}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func FindBook(c *gin.Context) {
	var book models.Book

	if error := models.DB.Where("id = ? ", c.Param("id")).First(&book).Error; error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context) {
	var book models.Book
	if error := models.DB.Where("id = ? ", c.Param("id")).First(&book).Error; error != nil {
		c.JSON(http.StatusBadRequest, "Record Not Found")
		return
	}

	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := models.DB.Model(&book).Updates(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}
