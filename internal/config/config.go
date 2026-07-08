package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL      string
	CS2LabPort       string
	OxeliaGatewayMode bool
	CS2LabStaticDir  string
}

func Load() *Config {
	godotenv.Load()

	return &Config{
		DatabaseURL:      getEnv("DATABASE_URL", "postgres://postgres:password@localhost:5432/oxelia51?sslmode=disable"),
		CS2LabPort:       getEnv("CS2LAB_PORT", "8001"),
		OxeliaGatewayMode: getEnv("OXELIA_GATEWAY_MODE", "false") == "true",
		CS2LabStaticDir:  getEnv("CS2LAB_STATIC_DIR", "./data"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
