package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	DB struct {
		DB_NAME     string `yaml : "db_name"`
		DB_USER     string `yaml : "db_user"`
		DB_PASSWORD string `yaml : "db_password"`
	} `yaml : "db"`
	Server struct {
		Port string `yaml : "port"`
	} `yaml : "server"`
}

func Load() (*Config, error) {
	configPath := getConfigPath("config/dev.yml")
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("cannot open config file")
	}
	defer file.Close()
	var cfg Config
	if err := yaml.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("failed to decode config")
	}

	return &cfg, nil
}

func getConfigPath(defaultPath string) string {
	_, b, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(b)

	projectRoot := filepath.Join(baseDir, "../..")

	fullPath := filepath.Join(projectRoot, defaultPath)
	if _, err := os.Stat(fullPath); err == nil {
		return fullPath
	}

	return defaultPath
}
