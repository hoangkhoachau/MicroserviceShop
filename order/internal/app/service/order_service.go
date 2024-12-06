package service

import "github.com/hoangkhoachau/MicroserviceShop/order/internal/app/dto"

type IOrderService interface {
	OrderCompleted(id string) error
	GetOrderByUserId(userId int) (*dto.GetOrderByUserIdResponse, error)
}
