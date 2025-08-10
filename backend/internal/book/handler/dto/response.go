package dto

type BookResponse struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Year      int    `json:"year,omitempty"`
}
