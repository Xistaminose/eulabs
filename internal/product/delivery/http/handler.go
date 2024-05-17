package http

import (
	"eulabs/internal/entity"
	"eulabs/internal/product/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productUsecase *usecase.ProductUsecase
}

func NewProductHandler(e *echo.Echo, u *usecase.ProductUsecase) {
	handler := &ProductHandler{u}

	e.POST("/products", handler.CreateProduct)
	e.GET("/products", handler.FetchProducts)
	e.GET("/products/:id", handler.GetProductByID)
	e.PUT("/products/:id", handler.UpdateProduct)
	e.DELETE("/products/:id", handler.DeleteProduct)
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	product := new(entity.Product)
	if err := c.Bind(product); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	product, err := h.productUsecase.CreateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) FetchProducts(c echo.Context) error {
	products, err := h.productUsecase.FetchProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProductByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := h.productUsecase.GetProductByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	product := new(entity.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	id, _ := strconv.Atoi(c.Param("id"))
	product.ID = id
	if err := h.productUsecase.UpdateProduct(product); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.productUsecase.DeleteProduct(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}
