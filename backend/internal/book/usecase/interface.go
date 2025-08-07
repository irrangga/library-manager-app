package usecase

import (
	"backend/internal/book/entity"
	"context"
)

type BookUsecase interface {
	GetBooks(ctx context.Context) ([]entity.Book, error)
	AddBook(ctx context.Context, input entity.Book) (entity.Book, error)
}
