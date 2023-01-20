package users

import (
	"context"

	"github.com/rnwxyz/wishlist-sewa/model"
)

type UserRepository interface {
	CreateUser(user *model.User, ctx context.Context) (*model.User, error)
	DefaultAccount(owner *model.Owner) error
	FindUserByEmail(email *string, ctx context.Context) (*model.User, error)
	FindUserByID(id *uint, ctx context.Context) (*model.User, error)
	FindUsers(page *model.Pagination, ctx context.Context) ([]model.User, int, error)
	UpdateOwner(owner *model.Owner, ctx context.Context) error
	DeleteUser(user *model.User, ctx context.Context) error
}
