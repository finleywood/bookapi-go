package controllers

import (
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, books)
}

func FindBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	res := models.DB.First(&book, id)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error})
		return
	}
	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, book)
}
