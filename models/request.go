package models

type CreateBookRequest struct {
	Title       string `json:"title" validate:"required"`
	Author      string `json:"author" validate:"required"`
	Description string `json:"description"`
}

type UpdateBookRequest struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}
