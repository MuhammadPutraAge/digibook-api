package book

type BookRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Genre       string `json:"genre" validate:"required"`
	Author      string `json:"author" validate:"required"`
}
