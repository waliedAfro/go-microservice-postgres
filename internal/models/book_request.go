package models

type CreateBookRequest struct {

	//book data
	Title  string `json:"title" validate:"required,min=3,max=150"`
	Author string `json:"author" validate:"required,min=3,max=100"`

	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type SearchBookRequest struct {

	//pagination
	Page  int `json:"page"`
	Limit int `json:"limit"`

	// search by Name ..
	SearchTerm string `json:"searchTerm"`

	//filter
	Author string `json:"author"`
	Title  string `json:"title"`

	//sort by

}
