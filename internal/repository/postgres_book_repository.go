package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"go-microservice-postgres/internal/models"
)

type PostgresBookRepository struct {
	db *pgxpool.Pool
}

func NewPostgresBookRepository(db *pgxpool.Pool) *PostgresBookRepository {
	return &PostgresBookRepository{
		db: db,
	}
}

func (r *PostgresBookRepository) FindAll() ([]models.BookModel, error) {

	rows, err := r.db.Query(
		context.Background(),
		`SELECT id, title, author
		 FROM books
		 ORDER BY id`,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.BookModel

	for rows.Next() {
		var book models.BookModel

		if err := rows.Scan(&book.ID, &book.Title, &book.Author); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, rows.Err()
}

func (r *PostgresBookRepository) FindByID(id int) (*models.BookModel, error) {

	book := &models.BookModel{}

	err := r.db.QueryRow(
		context.Background(),
		`SELECT id,title,author
		 FROM books
		 WHERE id=$1`, id).Scan(&book.ID, &book.Title, &book.Author)

	if err != nil {
		return nil, err
	}

	return book, nil
}

func (r *PostgresBookRepository) Create(book *models.BookModel) (*models.BookModel, error) {

	err := r.db.QueryRow(
		context.Background(),
		`INSERT INTO books(title,author)VALUES($1,$2)
		 RETURNING id`, book.Title, book.Author).Scan(&book.ID)

	if err != nil {
		return nil, err
	}

	return book, nil
}

func (r *PostgresBookRepository) Update(book *models.BookModel) (*models.BookModel, error) {

	_, err := r.db.Exec(
		context.Background(),
		`UPDATE books SET title=$1,author=$2 WHERE id=$3`,
		book.Title,
		book.Author,
		book.ID,
	)

	if err != nil {
		return nil, err
	}

	return book, nil
}

func (r *PostgresBookRepository) Delete(id int) error {

	_, err := r.db.Exec(
		context.Background(),
		`DELETE FROM books
		 WHERE id=$1`,
		id,
	)

	return err
}
