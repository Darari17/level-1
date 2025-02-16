package router

import (
	"level-1/controller"

	"github.com/gin-gonic/gin"
)

type BookRoutes struct {
	BookController controller.BookController
}

func NewBookRoutes(bc controller.BookController) *BookRoutes {
	return &BookRoutes{
		BookController: bc,
	}
}

func (r *BookRoutes) SetupRoutes(rg *gin.RouterGroup) {
	rg.POST("/books", r.BookController.Create)
	rg.PUT("/books/:id", r.BookController.Update)
	rg.DELETE("/books/:id", r.BookController.Delete)
	rg.GET("/books", r.BookController.Get)
	rg.GET("/books/:id", r.BookController.GetById)
}
