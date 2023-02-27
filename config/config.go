package config

import "github.com/techwithmat/kanbi/pkg/env"

type Config struct {
	Port     string
	Database *Database
}

type Database struct {
	Host     string
	Port     string
	User     string
	DB       string
	Password string
}

func NewConfig() *Config {
	env.LoadEnv()

	port := env.MustGet("PORT")

	if port == "" {
		port = "3000"
	}

	return &Config{
		Port: port,
		Database: &Database{
			Host:     env.MustGet("POSTGRES_HOST"),
			Port:     env.MustGet("POSTGRES_PORT"),
			User:     env.MustGet("POSTGRES_USER"),
			DB:       env.MustGet("POSTGRES_DB"),
			Password: env.MustGet("POSTGRES_PASSWORD"),
		},
	}
}
