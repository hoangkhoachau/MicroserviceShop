package factories

import (
	"github.com/hoangkhoachau/MicroserviceShop/catalog/config"
	"github.com/hoangkhoachau/MicroserviceShop/catalog/internal/app/repository"
	gormrepository "github.com/hoangkhoachau/MicroserviceShop/catalog/internal/infrastructure/repository/gorm-repository"
)

type RepositoryFactory struct {
	cfg               config.AppConfig
	connectionFactory ConnectionFactory
}

func NewRepositoryFactory(cfg config.AppConfig, connectionFactory ConnectionFactory) *RepositoryFactory {
	return &RepositoryFactory{cfg: cfg, connectionFactory: connectionFactory}
}

func (r *RepositoryFactory) GetProductRepository() repository.IProductRepository {
	switch r.cfg.System.DbManager {
	case "gorm":
		return gormrepository.NewGormProductRepository(r.connectionFactory.GetGormDb())
	}
	return nil
}
