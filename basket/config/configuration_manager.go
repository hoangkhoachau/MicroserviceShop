package config

import (
	"github.com/spf13/viper"
	"github.com/hoangkhoachau/MicroserviceShop/basket/pkg/constants"
)

func init() {
	viper.AddConfigPath(constants.ConfigPath)
	viper.SetConfigType(constants.ConfigType)
}

type configurationManager struct {
	appConfig    AppConfig
	queuesConfig QueuesConfig
}

func NewConfigurationManager() *configurationManager {
	return &configurationManager{
		appConfig:    readApplicationConfig(),
		queuesConfig: readQueuesConfig(),
	}
}

func (c *configurationManager) GetMysqlConfig() MongoConfig {
	return c.appConfig.Mongo
}

func (c *configurationManager) GetSystemConfig() SystemConfig {
	return c.appConfig.System
}

func (c *configurationManager) GetRabbitMqConfig() RabbitMqConfig {
	return c.appConfig.RabbitMq
}

func (c *configurationManager) GetAppConfig() AppConfig {
	return c.appConfig
}

func (c *configurationManager) GetQueuesConfig() QueuesConfig {
	return c.queuesConfig
}

func readApplicationConfig() AppConfig {
	viper.SetConfigName(constants.AppConfigName)
	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}
	var appConfig AppConfig
	if err := viper.Unmarshal(&appConfig); err != nil {
		panic(err.Error())
	}
	constants.TokenSecretKey = appConfig.System.TokenSecretKey
	return appConfig
}

func readQueuesConfig() QueuesConfig {
	viper.SetConfigName(constants.QueueConfigName)
	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}
	var queuesConfig QueuesConfig
	if err := viper.Unmarshal(&queuesConfig); err != nil {
		panic(err.Error())
	}
	return queuesConfig
}
