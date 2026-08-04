package dto

type UpdateBookRequest struct {

	//book data
	ID     int    `json:"id" validate:"required"`
	Title  string `json:"title" validate:"required,min=3,max=150"`
	Author string `json:"author" validate:"required,min=3,max=100"`
}
