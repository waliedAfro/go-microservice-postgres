package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// DATABASE_URL=postgres://postgres:password@localhost:5432/books?sslmode=disable
//DATABASE_URL=postgres://postgres:postgres@localhost:5432/postgres?ssl

type Config struct {
	DatabaseURL string
}

func Load() *Config {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	url := os.Getenv("GO_DATABASE_URL")

	if url == "" {
		log.Fatal("GO_DATABASE_URL is missing")
	}

	fmt.Println("Postgres URL loaded")

	return &Config{
		DatabaseURL: url,
	}
}
