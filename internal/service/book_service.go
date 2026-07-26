package service

import "go-microservice-postgres/internal/models"

type BookService interface {
	GetAllBooks() ([]models.BookModel, error)
	GetBookByID(id int) (*models.BookModel, error)
	CreateBook(book *models.BookModel) (*models.BookModel, error)
	UpdateBook(book *models.BookModel) (*models.BookModel, error)
	DeleteBook(id int) error
}
