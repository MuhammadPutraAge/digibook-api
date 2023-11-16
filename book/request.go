package book

type BookRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
	Author      string `json:"author"`
}
