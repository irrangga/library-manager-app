package mapper

import (
	"backend/internal/book/entity"
	"backend/internal/book/repo/model"
)

func ToBookEntities(bookModels []model.Book) []entity.Book {
	var books []entity.Book

	for _, model := range bookModels {
		books = append(books, ToBookEntity(model))
	}

	return books
}

func ToBookEntity(bookModel model.Book) entity.Book {
	return entity.Book{
		ID:        bookModel.ID,
		Title:     bookModel.Title,
		Author:    bookModel.Author,
		Publisher: bookModel.Publisher,
		Year:      bookModel.Year,
		CreatedAt: bookModel.CreatedAt,
		UpdatedAt: bookModel.UpdatedAt,
	}
}
