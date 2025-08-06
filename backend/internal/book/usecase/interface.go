package usecase

import (
	"backend/internal/book/entity"
	"context"
)

type BookUsecase interface {
	GetBooks(ctx context.Context) ([]entity.Book, error)
}
