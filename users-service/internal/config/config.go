package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		Logger   `yaml:"logger"`
		MongoDB  `yaml:"mongo_db"`
		SpawnApi `yaml:"spawn_api"`
	}

	// MongoDB =.
	MongoDB struct {
		URI      string `env-required:"true" yaml:"uri" env:"MONGO_URI"`
		User     string `env-required:"true" yaml:"user" env:"MONGO_USER"`
		Password string `env-required:"true" yaml:"password" env:"MONGO_PASSWORD"`
		Name     string `yaml:"name"`
	}

	// MongoDB =.
	SpawnApi struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}

	// Logger -.
	Logger struct {
		Level string `yaml:"log_level"`
	}
)

// NewConfig returns app config.
func NewConfig(path string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
