package grpc

import (
	"context"
	"testing"

	"github.com/FRFebi/template-service/domain"
	"github.com/FRFebi/template-service/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type BookUsecaseMock struct {
	mock.Mock
}

func (m *BookUsecaseMock) CreateBook(book *domain.Book) (*domain.Book, error) {
	args := m.Called(book)
	return args.Get(0).(*domain.Book), args.Error(1)
}

func (m *BookUsecaseMock) ShowBooks() ([]*domain.Book, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Book), args.Error(1)
}

func TestShowBooks(t *testing.T) {
	mockUsecase := new(BookUsecaseMock)
	handler := NewBookGRPC(mockUsecase)

	books := []*domain.Book{
		{Id: "1", Name: "Science"},
		{Id: "2", Name: "History"},
	}

	mockUsecase.On("ShowBooks").Return(books, nil)

	res, _ := handler.ShowBooks(context.Background(), &proto.EmptyRequest{})
	// assert.Nil(t, err)
	assert.NotNil(t, res)
}
