package products

import "github.com/labstack/echo/v4"

type ProductController interface {
	CreateProduct(c echo.Context) error
	GetProductByID(c echo.Context) error
	GetProducts(c echo.Context) error
	UpdateProduct(c echo.Context) error
	DeleteProduct(c echo.Context) error
}
