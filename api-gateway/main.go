package main

import (
	"github.com/hoangkhoachau/MicroserviceShop/api-gateway/api"
	"github.com/hoangkhoachau/MicroserviceShop/api-gateway/config"
	servicediscovery "github.com/hoangkhoachau/MicroserviceShop/api-gateway/service-discovery"
)

func main() {
	configManager := config.NewConfigManager()
	routesConfig := configManager.GetRoutesQunfig()
	if routesConfig.UseServiceDiscovery {

	}
	serviceDiscoveryFactory := servicediscovery.NewServiceDiscoveryFactory(routesConfig)

	apiFactory := api.NewApiFactory(routesConfig, *serviceDiscoveryFactory)
	apiFactory.GetApi().Start()
}
