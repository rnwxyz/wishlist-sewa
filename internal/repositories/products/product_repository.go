package products

import (
	"context"

	"github.com/rnwxyz/wishlist-sewa/model"
)

type ProductRepository interface {
	CreateProduct(product *model.Product, ctx context.Context) (*model.Product, error)
	FindProductByID(id *uint, ctx context.Context) (*model.Product, error)
	FindProducts(page *model.Pagination, ctx context.Context) ([]model.Product, int, error)
	UpdateProduct(product *model.Product, ctx context.Context) error
	DeleteProduct(product *model.Product, ctx context.Context) error
}
