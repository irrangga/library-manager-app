package usecase

import (
	"backend/generated/mockgen/book"
	"backend/internal/book/entity"
	"backend/internal/book/repo"
	"backend/pkg/errors"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_bookUsecase_GetBooks(t *testing.T) {
	ctrl := gomock.NewController(t)

	bookRepoMock := book.NewMockBookRepo(ctrl)

	books := []entity.Book{
		{
			ID:        1,
			Title:     "Book 1",
			Author:    "Author 1",
			Publisher: "Publisher 1",
			Year:      2021,
		},
		{
			ID:        2,
			Title:     "Book 2",
			Author:    "Author 2",
			Publisher: "Publisher 2",
			Year:      2022,
		},
	}

	type fields struct {
		bookRepo repo.BookRepo
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []entity.Book
		wantErr error
	}{
		{
			name: "GetBooks returns books successfully",
			fields: fields{
				bookRepo: bookRepoMock,
			},
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				bookRepoMock.EXPECT().GetBooks(gomock.Any()).Return(books, nil)
			},
			want:    books,
			wantErr: nil,
		},
		{
			name: "GetBooks returns error",
			fields: fields{
				bookRepo: bookRepoMock,
			},
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				bookRepoMock.EXPECT().GetBooks(gomock.Any()).Return([]entity.Book{}, errors.ErrInternalServerError)
			},
			want:    []entity.Book{},
			wantErr: errors.ErrInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &bookUsecase{
				bookRepo: tt.fields.bookRepo,
			}
			tt.mock()

			got, err := uc.GetBooks(tt.args.ctx)
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, tt.wantErr, err)
		})
	}
}

func Test_bookUsecase_AddBook(t *testing.T) {
	ctrl := gomock.NewController(t)

	bookRepoMock := book.NewMockBookRepo(ctrl)

	bookInput := entity.Book{
		Title:     "Book 1",
		Author:    "Author 1",
		Publisher: "Publisher 1",
		Year:      2021,
	}

	book := entity.Book{
		ID:        1,
		Title:     "Book 1",
		Author:    "Author 1",
		Publisher: "Publisher 1",
		Year:      2021,
	}

	type fields struct {
		bookRepo repo.BookRepo
	}
	type args struct {
		ctx   context.Context
		input entity.Book
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    entity.Book
		wantErr error
	}{
		{
			name: "AddBook returns book successfully",
			fields: fields{
				bookRepo: bookRepoMock,
			},
			args: args{
				ctx:   context.Background(),
				input: bookInput,
			},
			mock: func() {
				bookRepoMock.EXPECT().AddBook(gomock.Any(), bookInput).Return(book, nil)
			},
			want:    book,
			wantErr: nil,
		},
		{
			name: "AddBook returns error",
			fields: fields{
				bookRepo: bookRepoMock,
			},
			args: args{
				ctx:   context.Background(),
				input: bookInput,
			},
			mock: func() {
				bookRepoMock.EXPECT().AddBook(gomock.Any(), bookInput).Return(entity.Book{}, errors.ErrInternalServerError)
			},
			want:    entity.Book{},
			wantErr: errors.ErrInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &bookUsecase{
				bookRepo: tt.fields.bookRepo,
			}
			tt.mock()

			got, err := uc.AddBook(tt.args.ctx, tt.args.input)
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, tt.wantErr, err)
		})
	}
}

func Test_bookUsecase_GetBookByID(t *testing.T) {
	ctrl := gomock.NewController(t)

	bookRepoMock := book.NewMockBookRepo(ctrl)

	book := entity.Book{
		ID:        1,
		Title:     "Book 1",
		Author:    "Author 1",
		Publisher: "Publisher 1",
		Year:      2021,
	}

	type fields struct {
		bookRepo repo.BookRepo
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    entity.Book
		wantErr error
	}{
		{
			name: "GetBookByID returns book successfully",
			fields: fields{
				bookRepo: bookRepoMock,
			},
			args: args{
				ctx: context.Background(),
				id:  int64(1),
			},
			mock: func() {
				bookRepoMock.EXPECT().GetBookByID(gomock.Any(), int64(1)).Return(book, nil)
			},
			want:    book,
			wantErr: nil,
		},
		{
			name: "GetBookByID returns error",
			fields: fields{
				bookRepo: bookRepoMock,
			},
			args: args{
				ctx: context.Background(),
				id:  int64(1),
			},
			mock: func() {
				bookRepoMock.EXPECT().GetBookByID(gomock.Any(), int64(1)).Return(entity.Book{}, errors.ErrInternalServerError)
			},
			want:    entity.Book{},
			wantErr: errors.ErrInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &bookUsecase{
				bookRepo: tt.fields.bookRepo,
			}
			tt.mock()

			got, err := uc.GetBookByID(tt.args.ctx, tt.args.id)
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, tt.wantErr, err)
		})
	}
}

func Test_bookUsecase_UpdateBook(t *testing.T) {
	ctrl := gomock.NewController(t)

	bookRepoMock := book.NewMockBookRepo(ctrl)

	book := entity.Book{
		ID:        1,
		Title:     "Book 1",
		Author:    "Author 1",
		Publisher: "Publisher 1",
		Year:      2021,
	}

	type fields struct {
		bookRepo repo.BookRepo
	}
	type args struct {
		ctx   context.Context
		input entity.Book
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    entity.Book
		wantErr error
	}{
		{
			name: "UpdateBook returns book successfully",
			fields: fields{
				bookRepo: bookRepoMock,
			},
			args: args{
				ctx:   context.Background(),
				input: book,
			},
			mock: func() {
				bookRepoMock.EXPECT().UpdateBook(gomock.Any(), book).Return(book, nil)
			},
			want:    book,
			wantErr: nil,
		},
		{
			name: "UpdateBook returns error",
			fields: fields{
				bookRepo: bookRepoMock,
			},
			args: args{
				ctx:   context.Background(),
				input: book,
			},
			mock: func() {
				bookRepoMock.EXPECT().UpdateBook(gomock.Any(), book).Return(entity.Book{}, errors.ErrInternalServerError)
			},
			want:    entity.Book{},
			wantErr: errors.ErrInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &bookUsecase{
				bookRepo: tt.fields.bookRepo,
			}
			tt.mock()

			got, err := uc.UpdateBook(tt.args.ctx, tt.args.input)
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, tt.wantErr, err)
		})
	}
}
