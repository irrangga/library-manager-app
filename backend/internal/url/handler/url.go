package handler

import (
	"backend/internal/url/entity"
	"backend/internal/url/handler/dto"
	"backend/internal/url/usecase"
	"backend/pkg/constant"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type urlHandler struct {
	urlUsecase usecase.URLUsecase
}

func NewURLHandler(
	urlUsecase usecase.URLUsecase,
) *urlHandler {
	return &urlHandler{
		urlUsecase,
	}
}

// @Summary Cleanup URL
// @Description Cleanup URL
// @Tags URL
// @Accept json
// @Produce json
// @Param url body dto.URLRequest true "URL"
// @Success 200 {object} dto.URLResponse
// @Router /url [post]
func (h *urlHandler) CleanupURL(ctx *gin.Context) {
	var req dto.URLRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	switch req.Operation {
	case constant.URLOperationRedirection, constant.URLOperationCanonical, constant.URLOperationAll:
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("invalid operation").Error(),
		})
		return
	}

	input := entity.URL{
		InitialURL: req.URL,
		Operation:  req.Operation,
	}

	url, err := h.urlUsecase.CleanupURL(ctx.Request.Context(), input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"processed_url": url.ProcessedURL,
	})
}
