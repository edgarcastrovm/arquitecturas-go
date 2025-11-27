package http

import (
	"cashin-hexagonal/internal/application/usecase"
	"cashin-hexagonal/internal/domain/model"
	"cashin-hexagonal/internal/infrastructure/config"
	"cashin-hexagonal/internal/infrastructure/http/handler"
	"cashin-hexagonal/internal/infrastructure/http/middleware"
	p_postgres "cashin-hexagonal/internal/infrastructure/persistence/postgres"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(cfg *config.Config) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Auto migrate
	err = db.AutoMigrate(&model.CashIn{})
	if err != nil {
		fmt.Println("Error al ejecutar AutoMigrate:", err)
		return err
	}
	fmt.Println("Migración de tabla CashIn completada.")

	repo := p_postgres.NewCashInRepository(db)
	createUseCase := usecase.NewCreateCashInUseCase(repo)

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	api := r.Group("/api/v1")
	handler.RegisterHealthCheckRoutes(api)
	handler.RegisterCashInRoutes(api, createUseCase)

	return r.Run(":" + cfg.Port)
}
