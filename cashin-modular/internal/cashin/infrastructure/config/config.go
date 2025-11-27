package config

import "os"

type Config struct {
    Port        string
    DBHost      string
    DBPort      string
    DBUser      string
    DBPassword  string
    DBName      string
}

func Load() (*Config, error) {
    return &Config{
        Port:       getEnv("PORT", "8080"),
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     getEnv("DB_PORT", "5434"),
        DBUser:     getEnv("DB_USER", "postgres"),
        DBPassword: getEnv("DB_PASSWORD", "yondaime"),
        DBName:     getEnv("DB_NAME", "cash_db"),
    }, nil
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}
