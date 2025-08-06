package mapper

import (
	"backend/internal/book/entity"
	"backend/internal/book/handler/dto"
)

func ToBookResponses(books []entity.Book) []dto.BookResponse {
	var bookResponses []dto.BookResponse

	for _, book := range books {
		bookResponses = append(bookResponses, dto.BookResponse{
			ID:        book.ID,
			Title:     book.Title,
			Author:    book.Author,
			Publisher: book.Publisher,
			Year:      book.Year,
		})
	}

	return bookResponses
}
