package mapper

import (
	"backend/src/book/entity"
	"backend/src/book/handler/dto"
)

func ToBookResponses(books []entity.Book) []dto.BookResponse {
	var bookResponses []dto.BookResponse

	for _, book := range books {
		bookResponses = append(bookResponses, ToBookResponse(book))
	}

	return bookResponses
}

func ToBookResponse(book entity.Book) dto.BookResponse {
	return dto.BookResponse{
		ID:        book.ID,
		Title:     book.Title,
		Author:    book.Author,
		Publisher: book.Publisher,
		Year:      book.Year,
		ImageURL:  book.ImageURL,
	}
}
