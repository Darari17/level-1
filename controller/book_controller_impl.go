package controller

import (
	"fmt"
	"level-1/model/dto"
	"level-1/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookControllerImpl struct {
	BookService service.BookService
}

// Create implements BookController.
func (b *BookControllerImpl) Create(c *gin.Context) {
	var payload dto.BookRequest
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		fmt.Println("Error bind JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Request masuk:", payload)

	createBook, err := b.BookService.Create(payload)
	if err != nil {
		fmt.Println("Error dari Service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.WebResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   createBook,
	})
}

// Delete implements BookController.
func (b *BookControllerImpl) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = b.BookService.Delete(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	})
}

// Get implements BookController.
func (b *BookControllerImpl) Get(c *gin.Context) {
	books, err := b.BookService.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   books,
	})
}

// GetById implements BookController.
func (b *BookControllerImpl) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := b.BookService.GetById(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   book,
	})
}

// Update implements BookController.
func (b *BookControllerImpl) Update(c *gin.Context) {
	var payload dto.BookUpdateRequest

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateBook, err := b.BookService.Update(int64(id), payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   updateBook,
	})
}
