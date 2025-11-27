package main

import (
	"cashin-hexagonal/internal/infrastructure/config"
	"cashin-hexagonal/internal/infrastructure/http"
	"log"
)

func main() {
	errenv := config.LoadEnv()
	if errenv != nil {
		log.Fatalf("Error al cargar las variables de entorno: %v", errenv)
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	if err := http.StartServer(cfg); err != nil {
		log.Fatal("Server error:", err)
	}
}
