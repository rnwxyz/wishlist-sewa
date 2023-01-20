package resources

import "github.com/rnwxyz/wishlist-sewa/model"

type UserResource struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func (u *UserResource) FromModel(model *model.User) {
	u.ID = model.ID
	u.Name = model.Name
	u.Email = model.Email
	u.Role = model.Role
}
