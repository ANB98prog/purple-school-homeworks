package main

import (
	"fmt"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/configs"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/internal/handler"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/internal/service"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/pkg/jwt"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/pkg/logging"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/pkg/middlewares"
	"log"
	"net/http"
)

func main() {
	configs.ReadEnvironmentVariables()
	conf := configs.LoadConfig()
	config, err := logging.ReadLogConfig()
	if err != nil {
		panic(err)
	}
	logging.ConfigureLogrus(config)

	router := http.NewServeMux()

	// Repositories

	// Services
	authService := service.NewAuthCodeService()

	// JWT
	j := jwt.NewJWT(conf.Auth.Secret)

	// Handlers
	handler.NewAuthHandler(router, authService, j)
	handler.NewProductHandler(router, conf)

	stack := middlewares.Chain(middlewares.Logging)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	fmt.Println("Listening on port 8081")
	log.Fatal(server.ListenAndServe())
}
