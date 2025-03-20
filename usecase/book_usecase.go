package usecase

import (
	"github.com/FRFebi/template-service/domain"
)

type BookUsecase struct {
	BookRepository domain.BookRepository
}

func NewBookUsecase(repo domain.BookRepository) domain.BookUsecase {
	return &BookUsecase{BookRepository: repo}
}

func (u *BookUsecase) CreateBook(book *domain.Book) (*domain.Book, error) {
	return u.BookRepository.Create(book)
}

func (u *BookUsecase) ShowBooks() ([]*domain.Book, error) {
	return u.BookRepository.Show()
}
