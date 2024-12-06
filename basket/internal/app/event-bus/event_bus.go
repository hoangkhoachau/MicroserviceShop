package eventbus

import "github.com/hoangkhoachau/MicroserviceShop/basket/internal/app/event"

type IEventBus interface {
	Subscribe()
	PublishBasketVerifiedEvent(basketVerifiedEvent event.BasketVerified) error
}
