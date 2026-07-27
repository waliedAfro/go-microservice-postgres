package service

import (
	"go-microservice-postgres/internal/common/pagination"
	"go-microservice-postgres/internal/models"
)

type BookService interface {
	GetAllBooks(page, limit int) (*pagination.PaginationResponse[models.BookModel], error)
	GetBookByID(id int) (*models.BookModel, error)
	CreateBook(book *models.BookModel) (*models.BookModel, error)
	UpdateBook(book *models.BookModel) (*models.BookModel, error)
	DeleteBook(id int) error
}
