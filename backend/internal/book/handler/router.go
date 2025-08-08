package handler

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, h *bookHandler) {
	book := router.Group("/books")

	book.GET("", h.GetBooks)
	book.POST("", h.AddBook)
	book.GET("/:id", h.GetBookByID)
	book.PUT("/:id", h.UpdateBookByID)
}
