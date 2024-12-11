package repository

import (
	"github.com/hoangkhoachau/MicroserviceShop/catalog/internal/app/dto"
	"github.com/hoangkhoachau/MicroserviceShop/catalog/internal/app/model"
)

type IProductRepository interface {
	AddProduct(product *model.Product) error
	UpdateProduct(product *model.Product) error
	DeleteProductById(id int) error
	GetProducts(productFilter dto.ProductFilter) ([]*dto.ProductResponse, error)
	GetProductsByIdList(idList ...int) ([]*model.Product, error)
	UpdateProducts(products ...*model.Product) error
}
