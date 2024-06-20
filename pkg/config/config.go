package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
	} `json:"database"`
	HtmlTemplatePath string `json:"htmlTemplatePath"`
}

func Read() (Config, error) {
	var config Config
	// Для локальной разработки
	file, err := os.Open("../../config.json")
	// Для контейнера
	//file, err := os.Open("config.json")
	if err != nil {
		return config, fmt.Errorf("error opening config file: %w", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return config, fmt.Errorf("error decoding config file: %w", err)
	}

	return config, nil
}
