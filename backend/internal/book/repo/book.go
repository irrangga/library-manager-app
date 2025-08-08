package repo

import (
	"backend/internal/book/entity"
	"backend/internal/book/mapper"
	"backend/internal/book/repo/model"
	"context"

	"gorm.io/gorm"
)

type bookRepo struct {
	db *gorm.DB
}

func NewBookRepo(
	db *gorm.DB,
) BookRepo {
	return &bookRepo{
		db,
	}
}

func (r *bookRepo) GetBooks(ctx context.Context) ([]entity.Book, error) {
	var bookModels []model.Book

	err := r.db.WithContext(ctx).Find(&bookModels).Error
	if err != nil {
		return []entity.Book{}, err
	}

	return mapper.ToBookEntities(bookModels), nil
}

func (r *bookRepo) AddBook(ctx context.Context, book entity.Book) (entity.Book, error) {
	bookModel := model.Book{
		Title:     book.Title,
		Author:    book.Author,
		Publisher: book.Publisher,
		Year:      book.Year,
	}

	err := r.db.WithContext(ctx).Create(&bookModel).Error
	if err != nil {
		return entity.Book{}, err
	}

	return mapper.ToBookEntity(bookModel), nil
}

func (r *bookRepo) GetBookByID(ctx context.Context, id int64) (entity.Book, error) {
	var bookModel model.Book

	err := r.db.WithContext(ctx).First(&bookModel, id).Error
	if err != nil {
		return entity.Book{}, err
	}

	return mapper.ToBookEntity(bookModel), nil
}
