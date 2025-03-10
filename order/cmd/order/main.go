package main

import (
	"github.com/hoangkhoachau/MicroserviceShop/order/api"
	"github.com/hoangkhoachau/MicroserviceShop/order/config"
	_ "github.com/hoangkhoachau/MicroserviceShop/order/docs"
	"github.com/hoangkhoachau/MicroserviceShop/order/internal/factories"
	"github.com/hoangkhoachau/MicroserviceShop/order/registration"
)

// @title Order Service Api
// @version 1.0
// @description Order Service documents
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	configManager := config.NewConfigurationManager()
	appConfig := configManager.GetAppConfig()
	systemConfig := configManager.GetSystemConfig()
	queuesConfig := configManager.GetQueuesConfig()
	connectionFactory := factories.NewConnectionFactory(appConfig).ConnectDb().EventBusConnect()
	repositoryFactory := factories.NewRepositoryFactory(appConfig, *connectionFactory)
	eventBusFactory := factories.NewEventBusFactory(systemConfig, queuesConfig, *connectionFactory, *repositoryFactory)
	serviceFactory := factories.NewServiceFactory(*repositoryFactory, *eventBusFactory)
	registationFactory := registration.NewRegistrationFactory(appConfig)
	registationFactory.GetRegistrationService().Register()
	apiFactory := api.NewApiFactory(systemConfig, *serviceFactory)
	eventBusFactory.GetEventBus().Subscribe()
	apiFactory.GetApi().Start()
}
