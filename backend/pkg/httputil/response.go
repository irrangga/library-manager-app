package httputil

import (
	"backend/pkg/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, data any) {
	requestID, _ := ctx.Get(middleware.RequestIDKey)

	ctx.JSON(http.StatusOK, gin.H{
		"request_id": requestID,
		"data":       data,
	})
}

func Error(ctx *gin.Context, status int, message string) {
	requestID, _ := ctx.Get(middleware.RequestIDKey)

	ctx.JSON(status, gin.H{
		"request_id": requestID,
		"error":      message,
	})
}
