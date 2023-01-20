package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID             uint `gorm:"primarykey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	Name           string         `gorm:"type:varchar(255);not null"`
	Price          int            `gorm:"type:int;not null"`
	Address        string         `gorm:"type:varchar(255);not null"`
	LinkVideo      string         `gorm:"type:varchar(255);not null"`
	LinkGoogleMaps string         `gorm:"type:varchar(255);not null"`
	Length         int            `gorm:"type:int;not null"`
	Width          int            `gorm:"type:int;not null"`
	Pined          bool           `gorm:"type:bool;"`
	ProductType    string         `gorm:"type:enum('KONTRAKAN', 'RUKO', 'KOS');column:product_type"`
	Description    string
	KamarMandi     string
	KamarTidur     string
	Lantai         string
	Garasi         string
	Meja           string
	Kasur          string
	Ac             string
	KipasAangin    string
	RuangTamu      string
	Tv             string
	Wifi           string
	Image1         string
	Image2         string
	Image3         string
	Wishlist       []Wishlist
}

type ProductType string

const (
	ProductTypeKontrakan ProductType = "KONTRAKAN"
	ProductTypeRuko      ProductType = "RUKO"
	ProductTypeKos       ProductType = "KOS"
)
