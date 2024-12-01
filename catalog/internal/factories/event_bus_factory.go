package factories

import (
	"github.com/hoangkhoachau/MicroserviceShop/catalog/config"
	eventbus "github.com/hoangkhoachau/MicroserviceShop/catalog/internal/app/event-bus"
	"github.com/hoangkhoachau/MicroserviceShop/catalog/internal/infrastructure/event-bus/rabbitmq"
)

type eventBusFactory struct {
	systemConfig      config.SystemConfig
	queuesConfig      config.QueuesConfig
	connectionFactory ConnectionFactory
	serviceFactory    ServiceFactory
}

func NewEventBusFactory(
	systemConfig config.SystemConfig,
	queuesConfig config.QueuesConfig,
	connectionFactory ConnectionFactory,
	serviceFactory ServiceFactory,
) *eventBusFactory {
	return &eventBusFactory{
		systemConfig:      systemConfig,
		queuesConfig:      queuesConfig,
		connectionFactory: connectionFactory,
		serviceFactory:    serviceFactory,
	}
}

func (e *eventBusFactory) GetEventBus() eventbus.IEventBus {
	switch e.systemConfig.MessageBus {
	case "rabbitMq":
		return rabbitmq.NewRabbitMqEventBus(e.queuesConfig, e.connectionFactory.GetRabbitMqConnection(), e.serviceFactory.GetProductService())
	}
	return nil
}
