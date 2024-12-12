package main

import (
	_ "github.com/hoangkhoachau/MicroserviceShop/identity/docs"

	"github.com/hoangkhoachau/MicroserviceShop/identity/api"
	"github.com/hoangkhoachau/MicroserviceShop/identity/config"
	"github.com/hoangkhoachau/MicroserviceShop/identity/internal/factories"
	"github.com/hoangkhoachau/MicroserviceShop/identity/registration"
)

// @title Identity Service Api
// @version 1.0
// @description Identity Service documents
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	configurationManager := config.NewConfigurationManager()
	systemConfig := configurationManager.GetSystemConfig()
	appConfig := configurationManager.GetAppConfig()
	connectionFactory := factories.NewConnectionFactory(appConfig).DbConnect().MessageBusConnect()
	repositoryFactory := factories.NewRepositoryFactory(appConfig, *connectionFactory)
	serviceFactory := factories.NewServiceFactory(*connectionFactory, *repositoryFactory, appConfig)
	registrationFactory := registration.NewRegistrationFactory(appConfig)
	registrationFactory.GetRegistrationService().Register()
	apiFactory := api.NewApiFactory(systemConfig, *serviceFactory).GetApi()
	apiFactory.Start()
}
