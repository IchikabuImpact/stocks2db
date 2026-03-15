package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DB       DBConfig       `json:"db"`
	PriceAPI PriceAPIConfig `json:"priceApi"`
}

type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type PriceAPIConfig struct {
	BaseURL string `json:"baseUrl"`
}

func Load(path string) (*Config, error) {
	if path == "" {
		path = "config.json"
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve config path: %w", err)
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("config file not found: %s", absPath)
		}
		return nil, fmt.Errorf("failed to read config file %s: %w", absPath, err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file %s: %w", absPath, err)
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c *Config) validate() error {
	if c.DB.Host == "" || c.DB.Port == 0 || c.DB.Name == "" || c.DB.User == "" {
		return errors.New("db configuration is invalid: host, port, name, and user are required")
	}
	if c.PriceAPI.BaseURL == "" {
		return errors.New("priceApi.baseUrl is required")
	}
	return nil
}
