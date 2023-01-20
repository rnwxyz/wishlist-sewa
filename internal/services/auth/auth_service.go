package auth

import (
	"context"

	"github.com/rnwxyz/wishlist-sewa/dto/requests"
	"github.com/rnwxyz/wishlist-sewa/model"
)

type AuthService interface {
	Login(login *requests.Credential, ctx context.Context) (*model.Token, error)
	UserRegister(register *requests.UserRegister, ctx context.Context) error
	CreateDefaultAccount() error
}
