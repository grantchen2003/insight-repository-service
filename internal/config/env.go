package config

import (
	"errors"

	"github.com/joho/godotenv"
)

func LoadEnvVars(env string) (bool, error) {
	envIsProduction := false
	var err error = nil

	switch env {

	case "dev", "development":
		err = godotenv.Load(".env.dev")

	case "prod", "production":
		err = godotenv.Load(".env.prod")
		envIsProduction = true

	default:
		err = errors.New("no env vars found")
	}

	return envIsProduction, err
}
