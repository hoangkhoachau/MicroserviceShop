package repository

import (
	"context"

	"github.com/hoangkhoachau/MicroserviceShop/basket/internal/app/model"
)

type IBasketRepository interface {
	CreateBasket(ctx context.Context, basket model.Basket) error
	AddProductToBasket(ctx context.Context, basketId string, product model.Product) error
	GetBasketByUserId(ctx context.Context, userId int) (*model.Basket, error)
	EmptyBasket(ctx context.Context, userId int) error
}
