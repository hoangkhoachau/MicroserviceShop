package main

import (
	_ "github.com/hoangkhoachau/MicroserviceShop/catalog/docs"

	"github.com/hoangkhoachau/MicroserviceShop/catalog/api"
	"github.com/hoangkhoachau/MicroserviceShop/catalog/config"
	"github.com/hoangkhoachau/MicroserviceShop/catalog/internal/factories"
	"github.com/hoangkhoachau/MicroserviceShop/catalog/registration"
)

// @title Catalog Service Api
// @version 1.0
// @description Catalog Service documents
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	configurationManager := config.NewConfigurationManager()
	systemConfig := configurationManager.GetSystemConfig()
	queuesConfig := configurationManager.GetQueuesConfig()
	appConfig := configurationManager.GetAppConfig()

	connectionFactory := factories.NewConnectionFactory(appConfig).DbConnect().EventBusConnect()
	repositoryFactory := factories.NewRepositoryFactory(appConfig, *connectionFactory)
	serviceFactory := factories.NewServiceFactory(*repositoryFactory)
	eventBusFactory := factories.NewEventBusFactory(systemConfig, queuesConfig, *connectionFactory, *serviceFactory)
	registationFactory := registration.NewRegistrationFactory(appConfig)
	registationFactory.GetRegistrationService().Register()
	eventBusFactory.GetEventBus().Subscribe()
	apiFactory := api.NewApiFactory(systemConfig, *serviceFactory).GetApi()
	apiFactory.Start()
}
