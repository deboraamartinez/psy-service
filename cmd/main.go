package main

import (
	"log"
	"psy-service/config"
	"psy-service/routes"
)

func main() {
	config.LoadConfig()
	router := routes.SetupRouter()

	log.Fatal(router.Run(":8080"))
}
