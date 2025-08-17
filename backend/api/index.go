package handler

import (
	"backend/src/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router = initRouter()

func initRouter() *gin.Engine {
	r, _ := app.SetupRouter()
	return r
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
