package main

import (
	"log"

	"api/internal/infrastructure/restful"
	"api/internal/shared/config"
)

func main() {
	config, err := config.NewConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	restful.Start(config)
}
