package users

import (
	"context"

	"github.com/rnwxyz/wishlist-sewa/dto/resources"
)

type UserService interface {
	FindUserByID(id *uint, ctx context.Context) (*resources.UserResource, error)
}
