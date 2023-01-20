package requests

import "github.com/rnwxyz/wishlist-sewa/model"

type UserRegister struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Name     string `json:"name" validate:"required,personname"`
}

func (u *UserRegister) ToModel() *model.User {
	return &model.User{
		Email:    u.Email,
		Password: u.Password,
		Name:     u.Name,
	}
}
