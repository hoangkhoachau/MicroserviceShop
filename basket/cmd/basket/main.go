package main

import (
	"github.com/hoangkhoachau/MicroserviceShop/basket/api"
	"github.com/hoangkhoachau/MicroserviceShop/basket/config"
	_ "github.com/hoangkhoachau/MicroserviceShop/basket/docs"
	"github.com/hoangkhoachau/MicroserviceShop/basket/internal/factories"
)

// @title Basket Service Api
// @version 1.0
// @description Basket Service documents
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
	apiFactory := api.NewApiFactory(systemConfig, *serviceFactory)
	eventBusFactory.GetEventBus().Subscribe()
	apiFactory.GetApi().Start()
}
