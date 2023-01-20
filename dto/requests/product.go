package requests

import "github.com/rnwxyz/wishlist-sewa/model"

type ProductStore struct {
	Name           string `json:"name" validate:"required,name"`
	Price          int    `json:"price" validate:"required,number"`
	Address        string `json:"address" validate:"required"`
	ProductType    string `json:"product_type" validate:"required,productType"`
	LinkVideo      string `json:"link_video" validate:"required,url"`
	LinkGoogleMaps string `json:"link_google_maps" validate:"required,url"`
	Length         int    `json:"length" validate:"required,number"`
	Width          int    `json:"width" validate:"required,number"`
	Pined          bool   `json:"pined"`
	Description    string `json:"description"`
	KamarMandi     string `json:"kamar_mandi"`
	KamarTidur     string `json:"kamar_tidur"`
	Lantai         string `json:"lantai"`
	Garasi         string `json:"garasi"`
	Meja           string `json:"meja"`
	Kasur          string `json:"kasur"`
	Ac             string `json:"ac"`
	KipasAangin    string `json:"kipas_aangin"`
	RuangTamu      string `json:"ruang_tamu"`
	Tv             string `json:"tv"`
	Wifi           string `json:"wifi"`
	Image1         string `json:"image1" validate:"required,url"`
	Image2         string `json:"image2" validate:"omitempty,url"`
	Image3         string `json:"image3" validate:"omitempty,url"`
}

func (p *ProductStore) ToModel() *model.Product {
	return &model.Product{
		Name:           p.Name,
		Price:          p.Price,
		Address:        p.Address,
		ProductType:    p.ProductType,
		LinkVideo:      p.LinkVideo,
		LinkGoogleMaps: p.LinkGoogleMaps,
		Length:         p.Length,
		Width:          p.Width,
		Pined:          p.Pined,
		Description:    p.Description,
		KamarMandi:     p.KamarMandi,
		KamarTidur:     p.KamarTidur,
		Lantai:         p.Lantai,
		Garasi:         p.Garasi,
		Meja:           p.Meja,
		Kasur:          p.Kasur,
		Ac:             p.Ac,
		KipasAangin:    p.KipasAangin,
		RuangTamu:      p.RuangTamu,
		Tv:             p.Tv,
		Wifi:           p.Wifi,
		Image1:         p.Image1,
		Image2:         p.Image2,
		Image3:         p.Image3,
	}
}

type ProductUpdate struct {
	ID             uint
	Name           string `json:"name" validate:"omitempty,name"`
	Price          int    `json:"price" validate:"omitempty,number"`
	Address        string `json:"address"`
	ProductType    string `json:"product_type" validate:"omitempty,productType"`
	LinkVideo      string `json:"link_video" validate:"omitempty,url"`
	LinkGoogleMaps string `json:"link_google_maps" validate:"omitempty,url"`
	Length         int    `json:"length" validate:"omitempty,number"`
	Width          int    `json:"width" validate:"omitempty,number"`
	Pined          bool   `json:"pined"`
	Description    string `json:"description"`
	KamarMandi     string `json:"kamar_mandi"`
	KamarTidur     string `json:"kamar_tidur"`
	Lantai         string `json:"lantai"`
	Garasi         string `json:"garasi"`
	Meja           string `json:"meja"`
	Kasur          string `json:"kasur"`
	Ac             string `json:"ac"`
	KipasAangin    string `json:"kipas_aangin"`
	RuangTamu      string `json:"ruang_tamu"`
	Tv             string `json:"tv"`
	Wifi           string `json:"wifi"`
	Image1         string `json:"image1" validate:"omitempty,url"`
	Image2         string `json:"image2" validate:"omitempty,url"`
	Image3         string `json:"image3" validate:"omitempty,url"`
}

func (p *ProductUpdate) ToModel() *model.Product {
	return &model.Product{
		ID:             p.ID,
		Name:           p.Name,
		Price:          p.Price,
		Address:        p.Address,
		ProductType:    p.ProductType,
		LinkVideo:      p.LinkVideo,
		LinkGoogleMaps: p.LinkGoogleMaps,
		Length:         p.Length,
		Width:          p.Width,
		Pined:          p.Pined,
		Description:    p.Description,
		KamarMandi:     p.KamarMandi,
		KamarTidur:     p.KamarTidur,
		Lantai:         p.Lantai,
		Garasi:         p.Garasi,
		Meja:           p.Meja,
		Kasur:          p.Kasur,
		Ac:             p.Ac,
		KipasAangin:    p.KipasAangin,
		RuangTamu:      p.RuangTamu,
		Tv:             p.Tv,
		Wifi:           p.Wifi,
		Image1:         p.Image1,
		Image2:         p.Image2,
		Image3:         p.Image3,
	}
}
