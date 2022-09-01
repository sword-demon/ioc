package Config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"ioc/services"
	"log"
)

type ServiceConfig struct{}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{}
}

func (sc *ServiceConfig) OrderService() *services.OrderService {
	return services.NewOrderService()
}

func (sc ServiceConfig) DBService() *gorm.DB {
	db, err := gorm.Open("mysql",
		"root:1@tcp(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(10)
	return db
}
