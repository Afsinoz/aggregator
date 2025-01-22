package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = "/.gatorconfig.json"

func GetConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir + configFileName, nil
}

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg Config) SetUser(username string) {
	configDir, err := GetConfigDir()
	if err != nil {
		fmt.Println("Error of GetConfigDir", err)
	}

	cfg.CurrentUserName = username

	data, err := json.Marshal(cfg)
	if err != nil {
		fmt.Println("Error of marshaling struct", err)
	}
	if err := os.WriteFile(configDir, data, 0644); err != nil {
		fmt.Println("Writing file error")
	}

}

func Read() (Config, error) {
	var cfg Config
	configFileDir, err := GetConfigDir()
	if err != nil {
		return Config{}, err
	}
	data, err := os.ReadFile(configFileDir)
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("Unmarshal Error ", err)
	}
	return cfg, nil

}
