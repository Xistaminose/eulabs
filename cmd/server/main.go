package main

import (
	"eulabsapi/config"
	"eulabsapi/internal/product/delivery/http"
	"eulabsapi/internal/product/repository/sqlite"
	"eulabsapi/internal/product/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db := config.InitDB()
	productRepo := sqlite.NewProductRepository(db)
	productUsecase := usecase.NewProductUsecase(productRepo)
	http.NewProductHandler(e, productUsecase)

	e.Logger.Fatal(e.Start(":8080"))
}
