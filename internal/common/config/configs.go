package config

import (
	"os"
)

// Config is the configuration struct
type Config struct {
	DatabaseURL string `env:"DATABASE_URL"`
	ServerPort  string `env:"SERVER_PORT"`
}

// Instance is the global configuration
var Instance *Config

// Load loads the environment variables into the struct
func Load() {

	var c = Config{}

	c.DatabaseURL = os.Getenv("DATABASE_URL")
	c.ServerPort = os.Getenv("SERVER_PORT")

	Instance = &c
}
