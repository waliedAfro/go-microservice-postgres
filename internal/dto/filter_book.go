package dto

type BookFilter struct {
	// Pagination
	Page  int `json:"page"`
	Limit int `json:"limit"`

	// Search (searches Title and Author)
	Search string `json:"search"`

	// Individual filters
	Title  string `json:"title"`
	Author string `json:"author"`

	// Sorting
	// Supported values:
	// id, -id
	// title, -title
	// author, -author
	Sort string `json:"sort"`
}
