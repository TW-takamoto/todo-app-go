package config

import (
	"os"

	e "api/internal/shared/error"
)

type Config struct {
	DatabaseUrl string
}

func NewConfigFromEnv() (Config, error) {
	databaseUrl, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		return Config{}, e.ErrorBuilder(e.NotFound).Property("DATABASE_URL").Build()
	}
	return Config{DatabaseUrl: databaseUrl}, nil
}
