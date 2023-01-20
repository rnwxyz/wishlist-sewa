package products

import (
	"context"

	"github.com/rnwxyz/wishlist-sewa/dto/requests"
	"github.com/rnwxyz/wishlist-sewa/dto/resources"
	"github.com/rnwxyz/wishlist-sewa/internal/repositories/products"
	"github.com/rnwxyz/wishlist-sewa/model"
)

type productServieImpl struct {
	productRepository products.ProductRepository
}

// CreateProduct implements ProductService
func (s *productServieImpl) CreateProduct(product *requests.ProductStore, ctx context.Context) (uint, error) {
	productModel := product.ToModel()
	res, err := s.productRepository.CreateProduct(productModel, ctx)
	if err != nil {
		return 0, err
	}
	return res.ID, nil
}

// DeleteProduct implements ProductService
func (s *productServieImpl) DeleteProduct(id *uint, ctx context.Context) error {
	product := &model.Product{ID: *id}
	err := s.productRepository.DeleteProduct(product, ctx)
	if err != nil {
		return err
	}
	return nil
}

// FindProductByID implements ProductService
func (s *productServieImpl) FindProductByID(id *uint, ctx context.Context) (*resources.ProductDetailResource, error) {
	product, err := s.productRepository.FindProductByID(id, ctx)
	if err != nil {
		return nil, err
	}
	result := resources.ProductDetailResource{}
	result.FromModel(product)
	return &result, nil
}

// FindProducts implements ProductService
func (s *productServieImpl) FindProducts(page *model.Pagination, ctx context.Context) (*resources.ProductListResource, int, error) {
	product, count, err := s.productRepository.FindProducts(page, ctx)
	if err != nil {
		return nil, 0, err
	}
	var result resources.ProductListResource
	result.FromModel(product)
	return &result, count, nil
}

// UpdateProduct implements ProductService
func (s *productServieImpl) UpdateProduct(product *requests.ProductUpdate, ctx context.Context) error {
	productModel := product.ToModel()
	err := s.productRepository.UpdateProduct(productModel, ctx)
	if err != nil {
		return err
	}
	return nil
}

func NewProductService(productRepository products.ProductRepository) ProductService {
	return &productServieImpl{productRepository: productRepository}
}
