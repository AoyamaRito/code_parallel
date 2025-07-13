package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	APIKey string `json:"api_key"`
}

func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".make_parallel_config.json"), nil
}

func SetAPIKey(apiKey string) error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	config := Config{
		APIKey: apiKey,
	}

	data, err := json.Marshal(config)
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0600)
}

func GetAPIKey() (string, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return "", err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return "", err
	}

	return config.APIKey, nil
}