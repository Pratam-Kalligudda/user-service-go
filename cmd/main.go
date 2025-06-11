package main

import (
	"log"

	"github.com/Pratam-Kalligudda/user-service-go/config"
	"github.com/Pratam-Kalligudda/user-service-go/internal/api"
)

func main() {
	config, err := config.SetupConfig()
	if err != nil {
		log.Fatalf("error while setting up config ; %v ", err.Error())
	}
	api.SetupServer(config)
	return
}
