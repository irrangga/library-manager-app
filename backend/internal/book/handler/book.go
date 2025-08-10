package handler

import (
	"backend/internal/book/entity"
	"backend/internal/book/handler/dto"
	"backend/internal/book/mapper"
	"backend/internal/book/usecase"
	"backend/pkg/httputil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	bookUsecase usecase.BookUsecase
}

func NewBookHandler(
	bookUsecase usecase.BookUsecase,
) *bookHandler {
	return &bookHandler{
		bookUsecase,
	}
}

func (h *bookHandler) GetBooks(ctx *gin.Context) {
	books, err := h.bookUsecase.GetBooks(ctx.Request.Context())
	if err != nil {
		httputil.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	httputil.Success(ctx, mapper.ToBookResponses(books))
}

func (h *bookHandler) AddBook(ctx *gin.Context) {
	var req dto.BookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		httputil.Error(ctx, http.StatusBadRequest, err.Error())
	}

	input := entity.Book{
		Title:     req.Title,
		Author:    req.Author,
		Publisher: req.Publisher,
		Year:      req.Year,
	}

	book, err := h.bookUsecase.AddBook(ctx.Request.Context(), input)
	if err != nil {
		httputil.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	httputil.Success(ctx, mapper.ToBookResponse(book))
}

func (h *bookHandler) GetBookByID(ctx *gin.Context) {
	id, err := httputil.GetPathParamInt64(ctx, "id")
	if err != nil {
		httputil.Error(ctx, http.StatusBadRequest, err.Error())
	}

	book, err := h.bookUsecase.GetBookByID(ctx.Request.Context(), id)
	if err != nil {
		httputil.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	httputil.Success(ctx, mapper.ToBookResponse(book))
}

func (h *bookHandler) UpdateBookByID(ctx *gin.Context) {
	id, err := httputil.GetPathParamInt64(ctx, "id")
	if err != nil {
		httputil.Error(ctx, http.StatusBadRequest, err.Error())
	}

	var req dto.BookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		httputil.Error(ctx, http.StatusBadRequest, err.Error())
	}

	input := entity.Book{
		ID:        id,
		Title:     req.Title,
		Author:    req.Author,
		Publisher: req.Publisher,
		Year:      req.Year,
	}

	book, err := h.bookUsecase.UpdateBook(ctx.Request.Context(), input)
	if err != nil {
		httputil.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	httputil.Success(ctx, mapper.ToBookResponse(book))
}

func (h *bookHandler) DeleteBookByID(ctx *gin.Context) {
	id, err := httputil.GetPathParamInt64(ctx, "id")
	if err != nil {
		httputil.Error(ctx, http.StatusBadRequest, err.Error())
	}

	err = h.bookUsecase.DeleteBookByID(ctx.Request.Context(), id)
	if err != nil {
		httputil.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	httputil.Success(ctx, nil)
}
