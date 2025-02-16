package repository

import (
	"level-1/model/domain"

	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book domain.Book) (domain.Book, error)
	Update(book domain.Book) (domain.Book, error)
	Delete(id int64) error
	Get() ([]domain.Book, error)
	GetById(id int64) (domain.Book, error)
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &BookRepositoryImpl{
		db: db,
	}
}
