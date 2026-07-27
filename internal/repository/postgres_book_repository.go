package repository

import (
	"context"
	"math"

	"github.com/jackc/pgx/v5/pgxpool"

	"go-microservice-postgres/internal/common/pagination"
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

func (r *PostgresBookRepository) FindAll(page, limit int) (*pagination.PaginationResponse[models.BookModel], error) {

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	// Prevent clients from requesting huge pages
	if limit > 100 {
		limit = 100
	}

	offset := (page - 1) * limit

	// Get total number of books
	var total int

	err := r.db.QueryRow(
		context.Background(),
		`SELECT COUNT(*) FROM books`,
	).Scan(&total)

	if err != nil {
		return nil, err
	}

	// Retrieve current page
	rows, err := r.db.Query(
		context.Background(),
		`
		SELECT id, title, author
		FROM books
		ORDER BY id
		LIMIT $1 OFFSET $2
		`,
		limit,
		offset,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]models.BookModel, 0, limit)

	for rows.Next() {
		var book models.BookModel

		if err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
		); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	response := &pagination.PaginationResponse[models.BookModel]{
		Page:        page,
		Limit:       limit,
		Total:       total,
		TotalPages:  totalPages,
		HasNext:     page < totalPages,
		HasPrevious: page > 1,
		Data:        books,
	}

	return response, nil
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
