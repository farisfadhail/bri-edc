package main

import (
	"core-service/config"
	"core-service/internal/handler"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	mux := http.NewServeMux()
	mux.HandleFunc("/core/authorize", handler.AuthorizeHandler)

	server := &http.Server{
		Addr:    cfg.Addr(),
		Handler: mux,
	}

	log.Printf("Core-service running on port %s", cfg.Port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}
