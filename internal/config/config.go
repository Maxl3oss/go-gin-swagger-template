package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	DB_PSQL      string
	CONNECT_TYPE string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		Port:         os.Getenv("PORT"),
		DB_PSQL:      os.Getenv("DB_PSQL"),
		CONNECT_TYPE: os.Getenv("CONNECT_TYPE"),
	}, nil
}
