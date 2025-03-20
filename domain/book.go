package domain

import "time"

type Book struct {
	Id        string
	Name      string
	CreateAt  *time.Time
	UpdateAt  *time.Time
	DeletedAt *time.Time
}

type BookRepository interface {
	Create(book *Book) (*Book, error)
	Show() ([]*Book, error)
	Update(*Book) (*Book, error)
}

type BookUsecase interface {
	CreateBook(*Book) (*Book, error)
	ShowBooks() ([]*Book, error)
}
