package products

import (
	"context"

	"github.com/rnwxyz/wishlist-sewa/dto/requests"
	"github.com/rnwxyz/wishlist-sewa/dto/resources"
	"github.com/rnwxyz/wishlist-sewa/model"
)

type ProductService interface {
	FindProductByID(id *uint, ctx context.Context) (*resources.ProductDetailResource, error)
	FindProducts(page *model.Pagination, ctx context.Context) (*resources.ProductListResource, int, error)
	CreateProduct(product *requests.ProductStore, ctx context.Context) (uint, error)
	UpdateProduct(product *requests.ProductUpdate, ctx context.Context) error
	DeleteProduct(id *uint, ctx context.Context) error
}
