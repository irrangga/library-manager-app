package repo

import (
	"backend/internal/book/entity"
	"context"
)

type BookRepo interface {
	GetBooks(ctx context.Context) ([]entity.Book, error)
	AddBook(ctx context.Context, book entity.Book) (entity.Book, error)
	GetBookByID(ctx context.Context, id int64) (entity.Book, error)
	UpdateBook(ctx context.Context, book entity.Book) (entity.Book, error)
	DeleteBookByID(ctx context.Context, id int64) error
}
