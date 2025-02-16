package repository

import (
	"fmt"
	"level-1/model/domain"

	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
	db *gorm.DB
}

// Create implements BookRepository.
func (b *BookRepositoryImpl) Create(book domain.Book) (domain.Book, error) {
	fmt.Println("Menyimpan ke database:", book)

	err := b.db.Create(&book).Error
	if err != nil {
		fmt.Println("Error menyimpan data:", err)
	} else {
		fmt.Println("Data berhasil disimpan:", book)
	}
	return book, err
}

// Delete implements BookRepository.
func (b *BookRepositoryImpl) Delete(id int64) error {
	err := b.db.Where("id = ?", id).Delete(&domain.Book{}).Error
	return err
}

// Get implements BookRepository.
func (b *BookRepositoryImpl) Get() ([]domain.Book, error) {
	var books []domain.Book
	err := b.db.Find(&books).Error
	return books, err
}

// GetById implements BookRepository.
func (b *BookRepositoryImpl) GetById(id int64) (domain.Book, error) {
	var book domain.Book
	err := b.db.First(&book, "id = ?", id).Error
	return book, err
}

// Update implements BookRepository.
func (b *BookRepositoryImpl) Update(book domain.Book) (domain.Book, error) {
	err := b.db.First(&domain.Book{}, book.ID).Error
	if err != nil {
		return domain.Book{}, err
	}

	err = b.db.Model(&domain.Book{}).Where("id = ?", book.ID).Updates(domain.Book{
		Title:  book.Title,
		Author: book.Author,
		Year:   book.Year,
	}).Error
	if err != nil {
		return domain.Book{}, err
	}

	return book, err
}
