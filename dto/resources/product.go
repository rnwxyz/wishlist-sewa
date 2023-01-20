package resources

import "github.com/rnwxyz/wishlist-sewa/model"

type ProductResource struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Address     string `json:"address"`
	ProductType string `json:"product_type"`
	Image1      string `json:"image1"`
	Image2      string `json:"image2"`
	Image3      string `json:"image3"`
}

func (p *ProductResource) FromModel(model *model.Product) {
	p.ID = model.ID
	p.Name = model.Name
	p.Price = model.Price
	p.Address = model.Address
	p.ProductType = model.ProductType
	p.Image1 = model.Image1
	p.Image2 = model.Image2
	p.Image3 = model.Image3
}

type ProductListResource []ProductResource

func (p *ProductListResource) FromModel(model []model.Product) {
	for _, product := range model {
		var productResource ProductResource
		productResource.FromModel(&product)
		*p = append(*p, productResource)
	}
}

type ProductDetailResource struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Price          int    `json:"price"`
	Address        string `json:"address"`
	ProductType    string `json:"product_type"`
	LinkVideo      string `json:"link_video"`
	LinkGoogleMaps string `json:"link_google_maps"`
	Length         int    `json:"length"`
	Width          int    `json:"width"`
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
	Image1         string `json:"image1"`
	Image2         string `json:"image2"`
	Image3         string `json:"image3"`
}

func (p *ProductDetailResource) FromModel(model *model.Product) {
	p.ID = model.ID
	p.Name = model.Name
	p.Price = model.Price
	p.Address = model.Address
	p.ProductType = model.ProductType
	p.LinkVideo = model.LinkVideo
	p.LinkGoogleMaps = model.LinkGoogleMaps
	p.Length = model.Length
	p.Width = model.Width
	p.Pined = model.Pined
	p.Description = model.Description
	p.KamarMandi = model.KamarMandi
	p.KamarTidur = model.KamarTidur
	p.Lantai = model.Lantai
	p.Garasi = model.Garasi
	p.Meja = model.Meja
	p.Kasur = model.Kasur
	p.Ac = model.Ac
	p.KipasAangin = model.KipasAangin
	p.RuangTamu = model.RuangTamu
	p.Tv = model.Tv
	p.Wifi = model.Wifi
	p.Image1 = model.Image1
	p.Image2 = model.Image2
	p.Image3 = model.Image3
}
