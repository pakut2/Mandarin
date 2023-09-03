package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type EnvVariables struct {
	GO_ENV        string
	MONGO_URI     string
	DATABASE_NAME string
	PORT          string
}

var Env *EnvVariables

func LoadEnvVariables() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	if err := validateEnvVariables(); err != nil {
		return err
	}

	return nil
}

func validateEnvVariables() error {
	environment := os.Getenv("GO_ENV")
	if environment == "" {
		return errors.New("GO_ENV environment variable is not set")
	}

	mongoUri := os.Getenv("MONGO_URI")
	if mongoUri == "" {
		return errors.New("MONGO_URI environment variable is not set")
	}

	databaseName := os.Getenv("DATABASE_NAME")
	if databaseName == "" {
		return errors.New("DATABASE_NAME environment variable is not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		return errors.New("PORT environment variable is not set")
	}

	Env = &EnvVariables{
		GO_ENV:        environment,
		MONGO_URI:     mongoUri,
		DATABASE_NAME: databaseName,
		PORT:          port,
	}

	return nil
}
