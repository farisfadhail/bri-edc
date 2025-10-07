package main

import (
	"bri-edc/api/config"
	"bri-edc/api/internal/injector"
	"bri-edc/api/internal/routes"
	"fmt"
	"log"
)

func main() {
	ct, err := injector.InitializeApp()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	app := config.NewFiber()

	routes.SetupRouter(app, ct)

	port := config.GetEnv("PORT", "")
	err = app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err.Error())
	}
}
