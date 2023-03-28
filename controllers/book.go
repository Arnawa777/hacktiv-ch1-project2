package controllers

import (
	"net/http"
	"project-02/models"

	"github.com/gin-gonic/gin"
)

// Create books
func (c *Controllers) CreateBook(ctx *gin.Context) {
	var (
		request models.CreateBookRequest
	)

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book := models.Book{
		Title:       request.Title,
		Author:      request.Author,
		Description: request.Description,
	}

	if err := c.masterDB.Create(&book).Error; err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, book)
	return
}

// Get all books
func (c *Controllers) GetBooks(ctx *gin.Context) {
	books := []models.Book{}

	err := c.masterDB.Find(&books).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, books)
	return
}

// Get book by id
func (c *Controllers) GetBookById(ctx *gin.Context) {
	id := ctx.Param("id")
	// book := []models.Book{}
	var book models.Book

	err := c.masterDB.First(&book, id).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	ctx.JSON(http.StatusOK, book)
	return
}

// Update book
func (c *Controllers) UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	var book models.Book

	err := c.masterDB.First(&book, id).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	if err := ctx.BindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.masterDB.Save(&book)
	ctx.JSON(http.StatusOK, book)
	return
}

// Delete book
func (c *Controllers) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	book := []models.Book{}

	err := c.masterDB.First(&book, id).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.masterDB.Delete(&book, id)
	ctx.JSON(http.StatusOK, "Book deleted successfully")
	return
}
