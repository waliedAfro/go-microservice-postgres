package repository

import (
	"go-microservice-postgres/internal/common/pagination"
	"go-microservice-postgres/internal/dto"
	"go-microservice-postgres/internal/models"
)

type BookRepository interface {
	FindAll(filter *dto.BookFilter) (*pagination.PaginationResponse[models.BookModel], error)
	FindByID(id int) (*models.BookModel, error)
	Create(book *dto.CreateBookRequest) (*models.BookModel, error)
	Update(book *models.BookModel) (*models.BookModel, error)
	Delete(id int) error
}
