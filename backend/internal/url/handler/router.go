package handler

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, h *urlHandler) {
	url := router.Group("/url")

	url.POST("", h.CleanupURL)
}
