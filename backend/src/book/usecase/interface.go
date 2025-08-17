package usecase

import (
	"backend/src/book/entity"
	"context"
)

type BookUsecase interface {
	GetBooks(ctx context.Context) ([]entity.Book, error)
	AddBook(ctx context.Context, input entity.Book) (entity.Book, error)
	GetBookByID(ctx context.Context, id int64) (entity.Book, error)
	UpdateBook(ctx context.Context, input entity.Book) (entity.Book, error)
	DeleteBookByID(ctx context.Context, id int64) error
}
