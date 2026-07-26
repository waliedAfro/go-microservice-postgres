package service

import (
	"go-microservice-postgres/internal/models"
	"go-microservice-postgres/internal/repository"
)

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{
		repo: repo,
	}
}

func (s *bookService) GetAllBooks() ([]models.BookModel, error) {
	return s.repo.FindAll()
}

func (s *bookService) GetBookByID(id int) (*models.BookModel, error) {
	return s.repo.FindByID(id)
}

func (s *bookService) CreateBook(book *models.BookModel) (*models.BookModel, error) {
	return s.repo.Create(book)
}

func (s *bookService) UpdateBook(book *models.BookModel) (*models.BookModel, error) {
	return s.repo.Update(book)
}

func (s *bookService) DeleteBook(id int) error {
	return s.repo.Delete(id)
}
