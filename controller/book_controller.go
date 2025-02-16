package controller

import (
	"level-1/service"

	"github.com/gin-gonic/gin"
)

type BookController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	GetById(c *gin.Context)
	Get(c *gin.Context)
}

func NewBookController(bookService service.BookService) BookController {
	return &BookControllerImpl{
		BookService: bookService,
	}
}
