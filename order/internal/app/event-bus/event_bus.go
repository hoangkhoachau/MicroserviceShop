package eventbus

import "github.com/hoangkhoachau/MicroserviceShop/order/internal/app/event"

type IEventBus interface {
	Subscribe()
	PublishOrderCompletedEvent(orderCompleted event.OrderCompleted) error
}
