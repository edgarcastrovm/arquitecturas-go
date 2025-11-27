package kernel

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() (*Config, error) {
	return &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "clave_secreta"),
		DBName:     getEnv("DB_NAME", "db"),
	}, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func LoadEnv(envFile ...string) error {

	if len(envFile) > 0 {
		// Se habilita ambiente test
		fmt.Println("Variables entorno test:", envFile[0])
		if err := godotenv.Load(envFile[0]); err != nil {
			fmt.Println("No se encontró archivo: %v", envFile[0])
			return err
		}
	} else {
		// Si habilita ambiente real
		fmt.Println("Variables entorno real")
		// Cargar variables de entorno
		if err := godotenv.Load(); err != nil {
			fmt.Println("No se encontró archivo .env")
			return err
		}
	}
	return nil
}
