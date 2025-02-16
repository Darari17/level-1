package service

import (
	"level-1/model/dto"
	"level-1/repository"

	"github.com/go-playground/validator/v10"
)

type BookService interface {
	Create(book dto.BookRequest) (dto.BookResponse, error)
	Update(id int64, book dto.BookUpdateRequest) (dto.BookResponse, error)
	Delete(id int64) error
	GetById(id int64) (dto.BookResponse, error)
	Get() ([]dto.BookResponse, error)
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &BookServiceImpl{
		BookRepository: bookRepo,
		validate:       validator.New(),
	}
}
