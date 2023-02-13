package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		Grpc `yaml:"grpc"`
	}

	// Grpc -.
	Grpc struct {
		Port              string `yaml:"port"`
		MaxConnectionIdle int64  `yaml:"max_connection_idle"`
		Timeout           int64  `yaml:"timeout"`
		MaxConnectionAge  int64  `yaml:"max_connection_age"`
		Host              string `yaml:"host"`
	}
)

// NewConfig returns app config.
func NewConfig(path string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
