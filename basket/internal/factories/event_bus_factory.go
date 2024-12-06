package factories

import (
	"github.com/hoangkhoachau/MicroserviceShop/basket/config"
	eventbus "github.com/hoangkhoachau/MicroserviceShop/basket/internal/app/event-bus"
	rabbitMqEventBus "github.com/hoangkhoachau/MicroserviceShop/basket/internal/infrastructure/event-bus"
)

type EventBusFactory struct {
	systemConfig      config.SystemConfig
	queuesConfig      config.QueuesConfig
	connectionFactory ConnectionFactory
	repositoryFactory RepositoryFactory
}

func NewEventBusFactory(
	systemConfig config.SystemConfig,
	queuesConfig config.QueuesConfig,
	connectionFactory ConnectionFactory,
	repositoryFactory RepositoryFactory,
) *EventBusFactory {
	return &EventBusFactory{
		systemConfig:      systemConfig,
		queuesConfig:      queuesConfig,
		connectionFactory: connectionFactory,
		repositoryFactory: repositoryFactory,
	}
}

func (e *EventBusFactory) GetEventBus() eventbus.IEventBus {
	switch e.systemConfig.MessageBus {
	case "rabbitMq":
		return rabbitMqEventBus.NewRabbitMqEventBus(e.queuesConfig, e.connectionFactory.GetRabbitMqConnection(), e.repositoryFactory.GetBasketRepository())
	}
	return nil
}
