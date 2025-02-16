package service

import (
	"fmt"
	"level-1/model/domain"
	"level-1/model/dto"
	"level-1/repository"

	"github.com/go-playground/validator/v10"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
	validate       *validator.Validate
}

// Create implements BookService.
func (b *BookServiceImpl) Create(book dto.BookRequest) (dto.BookResponse, error) {
	fmt.Println("Memvalidasi data:", book)

	err := b.validate.Struct(book)
	if err != nil {
		fmt.Println("Validasi gagal:", err)
		return dto.BookResponse{}, err
	}

	request := domain.Book{
		Title:  book.Title,
		Author: book.Author,
		Year:   book.Year,
	}

	fmt.Println("Data yang akan disimpan:", request)

	createBook, err := b.BookRepository.Create(request)
	if err != nil {
		fmt.Println("Gagal menyimpan ke database:", err)
		return dto.BookResponse{}, err
	}

	fmt.Println("Data berhasil disimpan:", createBook)

	response := dto.BookResponse{
		ID:     createBook.ID,
		Title:  createBook.Title,
		Author: createBook.Author,
		Year:   createBook.Year,
	}

	return response, nil
}

// Delete implements BookService.
func (b *BookServiceImpl) Delete(id int64) error {
	book, err := b.BookRepository.GetById(id)
	if err != nil {
		return err
	}

	return b.BookRepository.Delete(book.ID)
}

// Get implements BookService.
func (b *BookServiceImpl) Get() ([]dto.BookResponse, error) {
	books, err := b.BookRepository.Get()
	if err != nil {
		return nil, err
	}

	var responses []dto.BookResponse
	for _, book := range books {
		responses = append(responses, dto.BookResponse{
			ID:     book.ID,
			Title:  book.Title,
			Author: book.Author,
			Year:   book.Year,
		})
	}

	return responses, nil
}

// GetById implements BookService.
func (b *BookServiceImpl) GetById(id int64) (dto.BookResponse, error) {
	book, err := b.BookRepository.GetById(id)
	if err != nil {
		return dto.BookResponse{}, err
	}

	response := dto.BookResponse{
		ID:     book.ID,
		Title:  book.Title,
		Author: book.Author,
		Year:   book.Year,
	}

	return response, nil
}

// Update implements BookService.
func (b *BookServiceImpl) Update(id int64, book dto.BookUpdateRequest) (dto.BookResponse, error) {
	if err := b.validate.Var(id, "gt=0"); err != nil {
		return dto.BookResponse{}, err
	}

	if err := b.validate.Struct(book); err != nil {
		return dto.BookResponse{}, err
	}

	currentBook, err := b.BookRepository.GetById(id)
	if err != nil {
		return dto.BookResponse{}, err
	}

	requestUpdate := domain.Book{
		ID:     id,
		Title:  book.Title,
		Author: book.Author,
		Year:   book.Year,
	}

	if book.Title == "" {
		requestUpdate.Title = currentBook.Title
	}
	if book.Author == "" {
		requestUpdate.Author = currentBook.Author
	}
	if book.Year == 0 {
		requestUpdate.Year = currentBook.Year
	}

	updateBook, err := b.BookRepository.Update(requestUpdate)
	if err != nil {
		return dto.BookResponse{}, err
	}

	response := dto.BookResponse{
		ID:     updateBook.ID,
		Title:  updateBook.Title,
		Author: updateBook.Author,
		Year:   updateBook.Year,
	}

	return response, nil
}
