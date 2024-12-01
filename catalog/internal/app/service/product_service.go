package service

import "github.com/hoangkhoachau/MicroserviceShop/catalog/internal/app/dto"

type IProductService interface {
	CreateProduct(createProductDto dto.CreateProductDTO) error
	GetProducts(productFilter dto.ProductFilter) ([]*dto.ProductResponse, error)
}
