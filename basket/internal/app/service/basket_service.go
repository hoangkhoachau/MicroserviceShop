package service

import (
	"github.com/hoangkhoachau/MicroserviceShop/basket/internal/app/dto"
	"github.com/hoangkhoachau/MicroserviceShop/basket/internal/app/event"
	"github.com/hoangkhoachau/MicroserviceShop/basket/internal/app/model"
)

type IBasketService interface {
	GetBasketByUserId(userId int) (*dto.GetBasketByUserIdResponse, error)
	AddProductToBasket(addProductRequest dto.AddProductToBasketRequest) error
	CreateBasket(basket model.Basket) error
	VerifyBasket(basketVerifiedEvent event.BasketVerified) error
}
