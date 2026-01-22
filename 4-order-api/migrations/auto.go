package main

import (
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/configs"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/internal/domain/entity"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/pkg/db"
)

func main() {
	configs.ReadEnvironmentVariables()
	conf := configs.LoadConfig()
	connection := db.NewDb(&conf.Db)

	if err := connection.AutoMigrate(&entity.User{}, &entity.Product{}, &entity.Order{}, &entity.OrderItem{}); err != nil {
		panic(err)
	}
}
