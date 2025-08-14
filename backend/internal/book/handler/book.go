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

// @Summary Get all books
// @Description Get all books
// @Tags Book
// @Accept json
// @Produce json
// @Success 200 {object} dto.BookResponse
// @Router /books [get]
func (h *bookHandler) GetBooks(ctx *gin.Context) {
	books, err := h.bookUsecase.GetBooks(ctx.Request.Context())
	if err != nil {
		httputil.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	httputil.Success(ctx, mapper.ToBookResponses(books))
}

// @Summary Add book
// @Description Add book
// @Tags Book
// @Accept json
// @Produce json
// @Param book body dto.BookRequest true "Book"
// @Success 200 {object} dto.BookResponse
// @Router /books [post]
func (h *bookHandler) AddBook(ctx *gin.Context) {
	var req dto.BookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		httputil.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	input := entity.Book{
		Title:     req.Title,
		Author:    req.Author,
		Publisher: req.Publisher,
		Year:      req.Year,
		ImageURL:  req.ImageURL,
	}

	book, err := h.bookUsecase.AddBook(ctx.Request.Context(), input)
	if err != nil {
		httputil.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	httputil.Success(ctx, mapper.ToBookResponse(book))
}

// @Summary Get book by id
// @Description Get book by id
// @Tags Book
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} dto.BookResponse
// @Router /books/{id} [get]
func (h *bookHandler) GetBookByID(ctx *gin.Context) {
	id, err := httputil.GetPathParamInt64(ctx, "id")
	if err != nil {
		httputil.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	book, err := h.bookUsecase.GetBookByID(ctx.Request.Context(), id)
	if err != nil {
		httputil.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	httputil.Success(ctx, mapper.ToBookResponse(book))
}

// @Summary Update book by id
// @Description Update book by id
// @Tags Book
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body dto.BookRequest true "Book"
// @Success 200 {object} dto.BookResponse
// @Router /books/{id} [put]
func (h *bookHandler) UpdateBookByID(ctx *gin.Context) {
	id, err := httputil.GetPathParamInt64(ctx, "id")
	if err != nil {
		httputil.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var req dto.BookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		httputil.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	input := entity.Book{
		ID:        id,
		Title:     req.Title,
		Author:    req.Author,
		Publisher: req.Publisher,
		Year:      req.Year,
		ImageURL:  req.ImageURL,
	}

	book, err := h.bookUsecase.UpdateBook(ctx.Request.Context(), input)
	if err != nil {
		httputil.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	httputil.Success(ctx, mapper.ToBookResponse(book))
}

// @Summary Delete book by id
// @Description Delete book by id
// @Tags Book
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} dto.BookResponse
// @Router /books/{id} [delete]
func (h *bookHandler) DeleteBookByID(ctx *gin.Context) {
	id, err := httputil.GetPathParamInt64(ctx, "id")
	if err != nil {
		httputil.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.bookUsecase.DeleteBookByID(ctx.Request.Context(), id)
	if err != nil {
		httputil.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	httputil.Success(ctx, nil)
}
