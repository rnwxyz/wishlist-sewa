package products

import (
	"context"
	"fmt"

	"github.com/rnwxyz/wishlist-sewa/model"
	"github.com/rnwxyz/wishlist-sewa/utils/myerrors"
	"gorm.io/gorm"
)

type productRepositoryImpl struct {
	db *gorm.DB
}

// CreateProduct implements ProductRepository
func (r *productRepositoryImpl) CreateProduct(product *model.Product, ctx context.Context) (*model.Product, error) {
	err := r.db.WithContext(ctx).Create(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

// DeleteProduct implements ProductRepository
func (r *productRepositoryImpl) DeleteProduct(product *model.Product, ctx context.Context) error {
	err := r.db.WithContext(ctx).Preload("Wishlists").First(product).Error
	if err != nil {
		return err
	}
	if len(product.Wishlist) > 0 {
		return myerrors.ErrRecordIsUsed
	}
	res := r.db.WithContext(ctx).Delete(product)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return myerrors.ErrRecordNotFound
	}
	return nil
}

// FindProductByID implements ProductRepository
func (r *productRepositoryImpl) FindProductByID(id *uint, ctx context.Context) (*model.Product, error) {
	product := model.Product{}
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// FindProducts implements ProductRepository
func (r *productRepositoryImpl) FindProducts(page *model.Pagination, ctx context.Context) ([]model.Product, int, error) {
	var products []model.Product
	var count int64

	productType := ctx.Value(model.ProductTypeKey)
	orderPrice := ctx.Value(model.OrderPriceKey)
	fmt.Println(productType, orderPrice)

	offset := (page.Limit * page.Page) - page.Limit
	query := r.db.WithContext(ctx).Model(model.Product{})
	if page.Q != "" {
		query = query.Where("name LIKE ?", "%"+page.Q+"%")
	}
	if productType.(string) != "" {
		query = query.Where("product_type = ?", productType)
	}
	if orderPrice.(string) != "" {
		query = query.Order("price " + orderPrice.(string))
	} else {
		query = query.Order("id DESC")
	}
	err := query.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Offset(offset).Limit(page.Limit).Find(&products).Error
	if err != nil {
		return nil, 0, err
	}
	return products, int(count), nil
}

// UpdateProduct implements ProductRepository
func (r *productRepositoryImpl) UpdateProduct(product *model.Product, ctx context.Context) error {
	res := r.db.WithContext(ctx).Model(product).Updates(product)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return myerrors.ErrRecordNotFound
	}
	return nil
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepositoryImpl{db}
}
