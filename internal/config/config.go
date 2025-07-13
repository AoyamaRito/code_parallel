package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	APIKey  string `json:"api_key"`
	Context string `json:"context"`
}

func getGlobalConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".make_parallel_config.json"), nil
}

func getLocalContextPath() string {
	return filepath.Join(".", ".make_parallel_context.json")
}

func loadGlobalConfig() (*Config, error) {
	configPath, err := getGlobalConfigPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{}, nil
		}
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func saveGlobalConfig(config *Config) error {
	configPath, err := getGlobalConfigPath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(config)
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0600)
}

type LocalContext struct {
	Context string `json:"context"`
}

func loadLocalContext() (*LocalContext, error) {
	contextPath := getLocalContextPath()
	
	data, err := os.ReadFile(contextPath)
	if err != nil {
		if os.IsNotExist(err) {
			return &LocalContext{}, nil
		}
		return nil, err
	}

	var context LocalContext
	if err := json.Unmarshal(data, &context); err != nil {
		return nil, err
	}

	return &context, nil
}

func saveLocalContext(context *LocalContext) error {
	contextPath := getLocalContextPath()
	
	data, err := json.Marshal(context)
	if err != nil {
		return err
	}

	return os.WriteFile(contextPath, data, 0644)
}

func SetAPIKey(apiKey string) error {
	config, err := loadGlobalConfig()
	if err != nil {
		return err
	}

	config.APIKey = apiKey
	return saveGlobalConfig(config)
}

func GetAPIKey() (string, error) {
	config, err := loadGlobalConfig()
	if err != nil {
		return "", err
	}
	return config.APIKey, nil
}

func SetContext(context string) error {
	localContext := &LocalContext{
		Context: context,
	}
	return saveLocalContext(localContext)
}

func GetContext() (string, error) {
	context, err := loadLocalContext()
	if err != nil {
		return "", err
	}
	return context.Context, nil
}