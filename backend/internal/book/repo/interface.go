package repo

import (
	"backend/internal/book/entity"
	"context"
)

type BookRepo interface {
	GetBooks(ctx context.Context) ([]entity.Book, error)
}
