package main

import (
	"github.com/hoangkhoachau/MicroserviceShop/email/config"
	"github.com/hoangkhoachau/MicroserviceShop/email/internal/factory"
)

func main() {
	configManager := config.NewConfigurationManager()
	eventBusFactory := factory.NewEventBusFactory(configManager.GetAppConfig())
	eventBusFactory.GetEventBus().Subscribe()
}
