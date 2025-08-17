package app

import (
	"backend/config"
	"backend/docs"
	"backend/pkg/middleware"
	bookhandler "backend/src/book/handler"
	bookrepo "backend/src/book/repo"
	bookusecase "backend/src/book/usecase"
	urlhandler "backend/src/url/handler"
	urlusecase "backend/src/url/usecase"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupRouter() (*gin.Engine, *config.Config) {
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
	urlUsecase := urlusecase.NewURLUsecase()

	// Handler initialization.
	bookHandler := bookhandler.NewBookHandler(bookUsecase)
	urlHandler := urlhandler.NewURLHandler(urlUsecase)

	// Router inititialization.
	router := gin.New()
	router.Use(middleware.Logger())

	bookhandler.RegisterRoutes(router, bookHandler)
	urlhandler.RegisterRoutes(router, urlHandler)

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginswagger.WrapHandler(swaggerfiles.Handler))

	return router, cfg
}
