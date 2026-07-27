package repository

import (
	"context"
	"fmt"
	"math"

	"github.com/jackc/pgx/v5/pgxpool"

	"go-microservice-postgres/internal/common/pagination"
	"go-microservice-postgres/internal/dto"
	"go-microservice-postgres/internal/models"
	"go-microservice-postgres/internal/repository/helper"
)

type PostgresBookRepository struct {
	db *pgxpool.Pool
}

func NewPostgresBookRepository(db *pgxpool.Pool) *PostgresBookRepository {
	return &PostgresBookRepository{
		db: db,
	}
}

func (r *PostgresBookRepository) FindAll(filter *dto.BookFilter) (*pagination.PaginationResponse[models.BookModel], error) {

	page, limit, offset := helper.BuildPagination(filter)

	// ------------------------------------------------------------------
	// Whitelist sorting columns
	// ------------------------------------------------------------------
	sortColumn, sortDirection := helper.BuildOrderBy(filter.Sort)

	baseQuery, args, index := helper.BuildWhereClause(filter)

	// ------------------------------------------------------------------
	// Count query
	// ------------------------------------------------------------------
	countQuery := "SELECT COUNT(*) " + baseQuery

	var total int

	err := r.db.QueryRow(context.Background(), countQuery, args...).Scan(&total)
	if err != nil {
		return nil, err
	}

	// ------------------------------------------------------------------
	// Data query
	// ------------------------------------------------------------------
	query := fmt.Sprintf(`
		SELECT id, title, author
		%s
		ORDER BY %s %s
		LIMIT $%d
		OFFSET $%d
	`,
		baseQuery,
		sortColumn,
		sortDirection,
		index,
		index+1,
	)

	args = append(args, filter.Limit, offset)

	rows, err := r.db.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]models.BookModel, 0, *limit)

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

	totalPages := 0
	if total > 0 {
		totalPages = int(math.Ceil(float64(total) / float64(*limit)))
	}

	return &pagination.PaginationResponse[models.BookModel]{
		Page:        *page,
		Limit:       *limit,
		Total:       total,
		TotalPages:  totalPages,
		HasNext:     *page < totalPages,
		HasPrevious: *page > 1,
		Data:        books,
	}, nil
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

func (r *PostgresBookRepository) Create(book *dto.CreateBookRequest) (*models.BookModel, error) {

	var bookId int

	err := r.db.QueryRow(
		context.Background(),
		`INSERT INTO books(title,author)VALUES($1,$2)
		 RETURNING id`, book.Title, book.Author).Scan(&bookId)

	if err != nil {
		return nil, err
	}

	return &models.BookModel{
		ID:     bookId,
		Title:  book.Title,
		Author: book.Author,
	}, nil
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
