package factories

import (
	"github.com/hoangkhoachau/MicroserviceShop/order/config"
	"github.com/hoangkhoachau/MicroserviceShop/order/internal/app/repository"
	mongoRepository "github.com/hoangkhoachau/MicroserviceShop/order/internal/infrastructure/repository/mongo-repository"
)

type RepositoryFactory struct {
	cfg               config.AppConfig
	connectionFactory ConnectionFactory
}

func NewRepositoryFactory(cfg config.AppConfig, connectionFactory ConnectionFactory) *RepositoryFactory {
	return &RepositoryFactory{cfg: cfg, connectionFactory: connectionFactory}
}

func (r *RepositoryFactory) GetOrderRepository() repository.IOrderRepository {
	switch r.cfg.System.DbDriver {
	case "mongo":
		return mongoRepository.NewMongoOrderRepository(r.connectionFactory.GetMongoDb().Collection("orders"))
	}
	return nil
}
