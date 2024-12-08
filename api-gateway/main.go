package main

import (
	"github.com/hoangkhoachau/MicroserviceShop/api-gateway/api"
	"github.com/hoangkhoachau/MicroserviceShop/api-gateway/config"
)

func main() {
	configManager := config.NewConfigManager()
	routesConfig := configManager.GetRoutesQunfig()
	apiFactory := api.NewApiFactory(routesConfig)
	apiFactory.GetApi().Start()
}
