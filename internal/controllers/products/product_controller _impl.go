package products

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rnwxyz/wishlist-sewa/dto/requests"
	"github.com/rnwxyz/wishlist-sewa/internal/services/products"
	"github.com/rnwxyz/wishlist-sewa/model"
)

type productControllerImpl struct {
	productService products.ProductService
}

// CreateProduct implements ProductController
func (d *productControllerImpl) CreateProduct(c echo.Context) error {
	var product requests.ProductStore
	if err := c.Bind(&product); err != nil {
		return err
	}
	if err := c.Validate(product); err != nil {
		return err
	}
	id, err := d.productService.CreateProduct(&product, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "new product success created",
		"data": echo.Map{
			"id": id,
		},
	})
}

// DeleteProduct implements ProductController
func (d *productControllerImpl) DeleteProduct(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	uintId := uint(id)
	if err := d.productService.DeleteProduct(&uintId, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete product",
	})
}

// GetProductByID implements ProductController
func (d *productControllerImpl) GetProductByID(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	uintId := uint(id)
	product, err := d.productService.FindProductByID(&uintId, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get product",
		"data":    product,
	})
}

// GetProducts implements ProductController
func (d *productControllerImpl) GetProducts(c echo.Context) error {
	var query model.Pagination
	query.NewPageQuery(c)
	ctx := c.Request().Context()
	ctx = context.WithValue(ctx, model.ProductTypeKey, c.QueryParam("product_type"))
	ctx = context.WithValue(ctx, model.OrderPriceKey, c.QueryParam("order_price"))

	products, count, err := d.productService.FindProducts(&query, ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get products",
		"data": echo.Map{
			"products": products,
			"count":    count,
			"page":     query.Page,
			"limit":    query.Limit,
		},
	})
}

// UpdateProduct implements ProductController
func (d *productControllerImpl) UpdateProduct(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	var product requests.ProductUpdate
	if err := c.Bind(&product); err != nil {
		return err
	}
	if err := c.Validate(product); err != nil {
		return err
	}
	product.ID = uint(id)
	if err := d.productService.UpdateProduct(&product, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update product",
	})
}

func NewProductController(productService products.ProductService) ProductController {
	return &productControllerImpl{
		productService: productService,
	}
}
