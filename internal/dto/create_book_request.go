package dto

type CreateBookRequest struct {

	//book data
	Title  string `json:"title" validate:"required,min=3,max=150"`
	Author string `json:"author" validate:"required,min=3,max=100"`
}
