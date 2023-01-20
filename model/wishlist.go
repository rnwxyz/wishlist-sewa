package model

import (
	"time"

	"gorm.io/gorm"
)

type Wishlist struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID    uint
	User      User
	ProductID uint
	Product   Product
}
