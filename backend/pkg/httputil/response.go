package httputil

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, data any) {
	response := gin.H{"data": data}
	ctx.JSON(http.StatusOK, response)
}

func Error(ctx *gin.Context, status int, message string) {
	ctx.JSON(status, gin.H{"error": message})
}
