package main

import (
	"cashin-modular/internal/cashin/application/usecase"
	http "cashin-modular/internal/cashin/infrastructure/http/handler"
	"cashin-modular/internal/cashin/infrastructure/persistence"
	"cashin-modular/internal/shared/kernel"
	"cashin-modular/internal/shared/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	errenv := kernel.LoadEnv()
	if errenv != nil {
		log.Fatalf("Error al cargar las variables de entorno: %v", errenv)
	}

	config, err := kernel.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load config:", err)

	}
	log.Println("Config loaded:", config)

	// Conexión DB
	dsn := "host=" + config.DBHost + " user=" + config.DBUser + " password=" + config.DBPassword + " dbname=" + config.DBName + " port=" + config.DBPort + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	// Repositorio
	repo := persistence.NewCashInRepository(db)

	// Use Cases
	createUC := usecase.NewCreateCashInUseCase(repo)
	getUC := usecase.NewGetCashInUseCase(repo)

	// Handler
	handler := http.NewCashInHandler(createUC, getUC)

	// Router
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	api := r.Group("/api/v1")
	api.GET("/hc", handler.GetHc)
	cashinGroup := api.Group("/cashin")
	{
		cashinGroup.POST("", handler.Create)
		cashinGroup.GET("/:id", handler.Get)
	}

	// Run server
	if err := r.Run(":" + config.ServerPort); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
