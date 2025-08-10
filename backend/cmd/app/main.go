package main

import (
	"backend/config"
	"backend/docs"
	bookhandler "backend/internal/book/handler"
	bookrepo "backend/internal/book/repo"
	bookusecase "backend/internal/book/usecase"
	"backend/pkg/middleware"
	"fmt"
	"log"

	swaggerfiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Configuration initialization.
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Database initialization.
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.Name,
		cfg.Postgres.Port,
		cfg.Postgres.SslMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Repo initialization.
	bookRepo := bookrepo.NewBookRepo(db)

	// Usecase initialization.
	bookUsecase := bookusecase.NewBookUsecase(bookRepo)

	// Handler initialization.
	bookHandler := bookhandler.NewBookHandler(bookUsecase)

	// Router inititialization.
	router := gin.New()
	router.Use(middleware.Logger())
	bookhandler.RegisterRoutes(router, bookHandler)

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginswagger.WrapHandler(swaggerfiles.Handler))

	httpPort := ":" + cfg.HTTP.Port
	router.Run(httpPort)
}
