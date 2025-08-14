package dto

type BookResponse struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher,omitempty"`
	Year      int    `json:"year,omitempty"`
	ImageURL  string `json:"image_url,omitempty"`
}
