package mapper

import (
	"backend/internal/book/entity"
	"backend/internal/book/repo/model"
)

func ToBookEntities(bookModels []model.Book) []entity.Book {
	var books []entity.Book

	for _, model := range bookModels {
		books = append(books, entity.Book{
			ID:        model.ID,
			Title:     model.Title,
			Author:    model.Author,
			Publisher: model.Publisher,
			Year:      model.Year,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		})
	}

	return books
}
