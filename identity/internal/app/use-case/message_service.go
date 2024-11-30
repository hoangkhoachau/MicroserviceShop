package usecase

import (
	"github.com/hoangkhoachau/MicroserviceShop/identity/internal/app/event"
)

type IMessageService interface {
	PublishUserCreatedEvent(event event.UserCreated) error
}
