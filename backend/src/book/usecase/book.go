package usecase

import (
	"backend/src/book/entity"
	"backend/src/book/repo"
	"context"
)

type bookUsecase struct {
	bookRepo repo.BookRepo
}

func NewBookUsecase(
	bookRepo repo.BookRepo,
) BookUsecase {
	return &bookUsecase{
		bookRepo,
	}
}

func (uc *bookUsecase) GetBooks(ctx context.Context) ([]entity.Book, error) {
	return uc.bookRepo.GetBooks(ctx)
}

func (uc *bookUsecase) AddBook(ctx context.Context, input entity.Book) (entity.Book, error) {
	return uc.bookRepo.AddBook(ctx, input)
}

func (uc *bookUsecase) GetBookByID(ctx context.Context, id int64) (entity.Book, error) {
	return uc.bookRepo.GetBookByID(ctx, id)
}

func (uc *bookUsecase) UpdateBook(ctx context.Context, input entity.Book) (entity.Book, error) {
	return uc.bookRepo.UpdateBook(ctx, input)
}

func (uc *bookUsecase) DeleteBookByID(ctx context.Context, id int64) error {
	return uc.bookRepo.DeleteBookByID(ctx, id)
}
