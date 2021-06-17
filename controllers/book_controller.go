package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/rest-api-gin-postgres-gorm/database"
	"github.com/isaqueveras/rest-api-gin-postgres-gorm/models"
)

// ShowBook Show a book
func ShowBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	db := database.GetDatabase()
	err = db.First(&book, newId).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot fin book: " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

// CreateBook Create a book
func CreateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})

		return
	}

	err = db.Create(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create book: " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

//  ShowBooks Show all books
func ShowBooks(c *gin.Context) {
	db := database.GetDatabase()

	var books []models.Book
	err := db.Find(&books).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list books: " + err.Error(),
		})

		return
	}

	c.JSON(200, books)
}

// UpdateBook Update a book
func UpdateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})

		return
	}

	err = db.Save(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update book: " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

// DeleteBook Delete a book
func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer" + err.Error(),
		})

		return
	}

	db := database.GetDatabase()
	err = db.Delete(&models.Book{}, newId).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete book: " + err.Error(),
		})

		return
	}

	c.Status(204)
}
