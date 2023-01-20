package users

import (
	"context"
	"strings"

	"github.com/rnwxyz/wishlist-sewa/model"
	"github.com/rnwxyz/wishlist-sewa/utils/myerrors"
	"gorm.io/gorm"
)

type UserRespositoryImpl struct {
	db *gorm.DB
}

// DefaultAccount implements UserRepository
func (r *UserRespositoryImpl) DefaultAccount(owner *model.Owner) error {
	err := r.db.First(&model.User{}).Error
	if err == nil {
		return nil
	}
	// user := model.User{
	// 	Email:    owner.User.Email,
	// 	Password: owner.User.Password,
	// 	Name:     owner.User.Name,
	// 	Role:     owner.User.Role,
	// }
	// err = r.db.Create(&user).Error
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(user)
	// owner.UserID = user.ID
	err = r.db.Create(owner).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateUser implements UserRepository
func (r *UserRespositoryImpl) CreateUser(user *model.User, ctx context.Context) (*model.User, error) {
	err := r.db.WithContext(ctx).Create(user).Error
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062") {
			return nil, myerrors.ErrEmailAlredyExist
		}
		return nil, err
	}
	return user, nil
}

// DeleteUser implements UserRepository
func (r *UserRespositoryImpl) DeleteUser(user *model.User, ctx context.Context) error {
	res := r.db.WithContext(ctx).Delete(user)
	err := res.Error
	if err != nil {
		return err
	}
	if res.RowsAffected == 0 {
		return myerrors.ErrRecordNotFound
	}
	return nil
}

// FindUserByEmail implements UserRepository
func (r *UserRespositoryImpl) FindUserByEmail(email *string, ctx context.Context) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("email = ?", *email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindUserByID implements UserRepository
func (r *UserRespositoryImpl) FindUserByID(id *uint, ctx context.Context) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("id = ?", *id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindUsers implements UserRepository
func (r *UserRespositoryImpl) FindUsers(page *model.Pagination, ctx context.Context) ([]model.User, int, error) {
	var users []model.User
	var count int64
	offset := (page.Limit * page.Page) - page.Limit

	if page.Q != "" {
		query := r.db.WithContext(ctx).Model(&model.User{}).
			Where("(name LIKE ? OR email LIKE ?)", "%"+page.Q+"%", "%"+page.Q+"%").
			Offset(offset).
			Limit(page.Limit).
			Order("id DESC")
		err := query.Find(&users).Error
		if err != nil {
			return nil, 0, err
		}

		err = query.Count(&count).Error
		if err != nil {
			return nil, 0, err
		}

		return users, int(count), nil
	}

	err := r.db.WithContext(ctx).Model(&model.User{}).
		Offset(offset).
		Limit(page.Limit).
		Order("id DESC").
		Find(&users).
		Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Model(&model.User{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	return users, int(count), nil
}

// UpdateOwner implements UserRepository
func (r *UserRespositoryImpl) UpdateOwner(owner *model.Owner, ctx context.Context) error {
	err := r.db.WithContext(ctx).Save(owner).Error
	if err != nil {
		return err
	}
	return nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRespositoryImpl{db: db}
}
