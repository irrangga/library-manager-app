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
