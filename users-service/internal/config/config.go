package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		Logger   `yaml:"logger"`
		MongoDB  `yaml:"mongo_db"`
		SpawnApi `yaml:"spawn_api"`
		Http     `yaml:"http"`
	}

	// MongoDB -.
	MongoDB struct {
		URI      string `env-required:"true" yaml:"uri" env:"MONGO_URI"`
		User     string `env-required:"true" yaml:"user" env:"MONGO_USER"`
		Password string `env-required:"true" yaml:"password" env:"MONGO_PASSWORD"`
		Name     string `yaml:"name"`
	}

	// SpawnApi -.
	SpawnApi struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}

	// Http -.
	Http struct {
		Host                   string `yaml:"host"`
		Port                   string `yaml:"port"`
		DefaultReadTimeout     int    `yaml:"default_read_timeout"`
		DefaultWriteTimeout    int    `yaml:"default_write_timeout"`
		DefaultShutdownTimeout int    `yaml:"default_shutdown_timeout"`
	}

	// Logger -.
	Logger struct {
		Level string `yaml:"log_level"`
	}
)

// NewConfig returns app config.
func NewConfig(path, envPath string) (*Config, error) {
	cfg := &Config{}

	err := setEnv(envPath)
	if err != nil {
		return nil, fmt.Errorf("config - setEnv: %w", err)
	}

	err = cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, fmt.Errorf("config - ReadConfig: %w", err)
	}

	return cfg, nil
}

func setEnv(path string) error {
	var file *os.File
	var err error
	file, err = os.Open(path)
	if err != nil {
		// this is need for tests files
		if file, err = os.Open("../../env.example"); err != nil {
			if file, err = os.Open("../../../../env.example"); err != nil {
				return fmt.Errorf("setEnv - Open: %w", err)
			}
		}
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				os.Setenv(key, value)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("setEnv - Scan: %w", err)
	}
	return nil
}
