package dto

type BookRequest struct {
	Title  string `json:"title" validate:"required,max=255"`
	Author string `json:"author" validate:"required,max=255"`
	Year   int    `json:"year" validate:"required"`
}

type BookUpdateRequest struct {
	Title  string `json:"title" validate:"omitempty,max=255"`
	Author string `json:"author" validate:"omitempty,max=255"`
	Year   int    `json:"year" validate:"omitempty"`
}

type BookResponse struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

type WebResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data,omitempty"`
}
