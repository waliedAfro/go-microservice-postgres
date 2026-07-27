package helper

import (
	"fmt"
	"go-microservice-postgres/internal/dto"
)

func BuildWhereClause(filter *dto.BookFilter) (string, []any, int) {

	baseQuery := `
		FROM books
		WHERE 1=1
	`

	args := make([]interface{}, 0)
	index := 1

	if filter.Title != "" {
		baseQuery += fmt.Sprintf(" AND title ILIKE $%d", index)
		args = append(args, "%"+filter.Title+"%")
		index++
	}

	if filter.Author != "" {
		baseQuery += fmt.Sprintf(" AND author ILIKE $%d", index)
		args = append(args, "%"+filter.Author+"%")
		index++
	}

	if filter.Search != "" {
		searchPattern := "%" + filter.Search + "%"
		baseQuery += fmt.Sprintf(" AND (title ILIKE $%d OR author ILIKE $%d)", index, index)
		args = append(args, searchPattern)
		index++
	}

	return baseQuery, args, index

}
func BuildOrderBy(sort string) (string, string) {

	sortColumn := "id"
	sortDirection := "ASC"

	switch sort {
	case "title":
		sortColumn = "title"
	case "-title":
		sortColumn = "title"
		sortDirection = "DESC"
	case "author":
		sortColumn = "author"
	case "-author":
		sortColumn = "author"
		sortDirection = "DESC"
	case "-id":
		sortDirection = "DESC"
	}
	return sortColumn, sortDirection
}

func BuildPagination(filter *dto.BookFilter) (*int, *int, *int) {

	// Default pagination
	if filter.Page <= 0 {
		filter.Page = 1
	}

	if filter.Limit <= 0 {
		filter.Limit = 10
	}

	if filter.Limit > 100 {
		filter.Limit = 100
	}

	offset := (filter.Page - 1) * filter.Limit

	return &filter.Page, &filter.Limit, &offset
}
