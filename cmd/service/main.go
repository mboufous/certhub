package main

import (
	"log"

	"github.com/mboufous/certhub/api"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	apiConfig := api.ApiConfig{
		Port: "8080",
		Host: "0.0.0.0",
	}
	router := api.SetupRoutes()

	api := api.New(router, &apiConfig)
	api.Start()

	return nil
}
