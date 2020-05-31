package config

import (
	"errors"
	"os"

	"github.com/jinzhu/configor"
)

// Config is a structure for application server and database settings
type Config struct {
	Server struct {
		// Type of the server could for now only grpc is accepted
		Type string
		Host string `env:"SERVER_HOST"`
		Port uint   `env:"SERVER_PORT"`
	}
	Database struct {
		// Type of the datbase for now only mongo is accepted
		Type     string
		Host     string `env:"DB_HOST"`
		Port     uint   `env:"DB_PORT"`
		User     string `env:"DB_USER"`
		Password string `env:"DB_PASSWORD"`
	}
}

// Load is a function to return config of the application
func Load(path string) (*Config, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, errors.New("file does not exists")
	}
	config := &Config{}
	err := configor.Load(config, path)
	if err != nil {
		return nil, err
	}

	return config, nil
}
