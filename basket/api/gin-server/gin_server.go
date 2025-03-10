package ginserver

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/hoangkhoachau/MicroserviceShop/basket/api/gin-server/middleware"
	"github.com/hoangkhoachau/MicroserviceShop/basket/config"
	"github.com/hoangkhoachau/MicroserviceShop/basket/internal/factories"
)

type GinServer struct {
	cfg            config.SystemConfig
	logger         *logrus.Logger
	router         *gin.Engine
	serviceFactory factories.ServiceFactory
}

func NewGinServer(serviceFactory factories.ServiceFactory, cfg config.SystemConfig) *GinServer {
	return &GinServer{serviceFactory: serviceFactory, cfg: cfg}
}

func (g *GinServer) Start() {
	g.create().
		generateSwagger().
		generateLogger().
		generateBasketGroup().
		listen()
}

func (g *GinServer) create() *GinServer {
	g.router = gin.Default()
	return g
}
func (g *GinServer) generateLogger() *GinServer {
	g.logger = logrus.New()
	g.logger.SetFormatter(&logrus.JSONFormatter{})
	file, _ := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	g.logger.SetOutput(file)
	return g

}
func (g *GinServer) generateSwagger() *GinServer {
	g.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return g
}

func (g *GinServer) generateBasketGroup() *GinServer {
	basketApi := NewBasketApi(g.serviceFactory.GetBasketService(), g.logger)

	g.router.Any("/check", func(ctx *gin.Context) {
		ctx.Writer.WriteHeader(http.StatusOK)
	})

	routerGroup := g.router.Group("baskets")

	routerGroup.GET("/:userId", middleware.Authorization("customer", "admin"), basketApi.GetBasketByUserID)
	routerGroup.POST("/addProduct", middleware.Authorization("customer"), basketApi.AddProductToBasket)
	routerGroup.PUT("/:userId/verify", middleware.Authorization("customer"), basketApi.VerifyBasketByUserId)
	return g
}

func (g *GinServer) listen() {
	address := fmt.Sprintf(":%d", g.cfg.Port)
	g.router.Run(address)
}
