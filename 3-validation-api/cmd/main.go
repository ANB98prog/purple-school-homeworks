package main

import (
	"fmt"
	"github.com/ANB98prog/purple-school-homeworks/3-validation-api/configs"
	"github.com/ANB98prog/purple-school-homeworks/3-validation-api/internal/verify"
	"log"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()

	emailVerificationService := verify.NewEmailVerificationService(conf.EmailSender)

	verify.NewVerifyHandler(router, verify.VerifyHandlerDeps{
		Config:              conf,
		VerificationService: emailVerificationService,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Listening on port 8081")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Could not start server. Error: %e", err)
	}
}
