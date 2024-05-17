package main

import (
	"eulabs/config"
	"eulabs/internal/product/delivery/http"
	"eulabs/internal/product/repository"
	"eulabs/internal/product/usecase"
	"log"
	"log/slog"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db, dbType, err := config.InitDB()
	if err != nil {
		slog.Error("failed to initialize database", err)
		log.Fatal(err)
	}

	productRepository, err := repository.NewProductRepository(db, dbType)
	if err != nil {
		slog.Error("failed to initialize product repository", err)
		log.Fatal(err)
	}

	productUsecase := usecase.NewProductUsecase(productRepository)
	http.NewProductHandler(e, productUsecase)

	e.Logger.Fatal(e.Start(":8080"))
}
