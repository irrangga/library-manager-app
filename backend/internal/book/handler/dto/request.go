package dto

type BookRequest struct {
	Title     string `json:"title" binding:"required"`
	Author    string `json:"author" binding:"required"`
	Publisher string `json:"publisher"`
	Year      int    `json:"year"`
}
