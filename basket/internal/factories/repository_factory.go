package factories

import (
	"github.com/hoangkhoachau/MicroserviceShop/basket/config"
	"github.com/hoangkhoachau/MicroserviceShop/basket/internal/app/repository"
	mongoRepository "github.com/hoangkhoachau/MicroserviceShop/basket/internal/infrastructure/repository/mongo-repository"
)

type RepositoryFactory struct {
	cfg               config.AppConfig
	connectionFactory ConnectionFactory
}

func NewRepositoryFactory(cfg config.AppConfig, connectionFactory ConnectionFactory) *RepositoryFactory {
	return &RepositoryFactory{cfg: cfg, connectionFactory: connectionFactory}
}

func (r *RepositoryFactory) GetBasketRepository() repository.IBasketRepository {
	switch r.cfg.System.DbDriver {
	case "mongo":
		return mongoRepository.NewMongoBasketRepository(r.connectionFactory.GetMongoDb().Collection("baskets"))
	}
	return nil
}
