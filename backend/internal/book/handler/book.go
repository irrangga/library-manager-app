package handler

import (
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
	}

	httputil.Success(ctx, mapper.ToBookResponses(books))
}
