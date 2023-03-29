package models

type CreateBookRequest struct {
	Title  string `json:"name_book" validate:"required"`
	Author string `json:"author" validate:"required"`
}

type UpdateBookRequest struct {
	Title  string `json:"name_book"`
	Author string `json:"author"`
}
