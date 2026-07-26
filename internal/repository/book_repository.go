package repository

import "go-microservice-postgres/internal/models"

type BookRepository interface {
	FindAll() ([]models.BookModel, error)
	FindByID(id int) (*models.BookModel, error)
	Create(book *models.BookModel) (*models.BookModel, error)
	Update(book *models.BookModel) (*models.BookModel, error)
	Delete(id int) error
}
