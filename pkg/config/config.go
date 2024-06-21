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
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		//configPath = "../../config.json" // Для локальной разработки
		configPath = "config.json" // путь по умолчанию для контейнера
	}

	file, err := os.Open(configPath)
	if err != nil {
		return config, fmt.Errorf("error opening config file: %w", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return config, fmt.Errorf("error decoding config file: %w", err)
	}

	// Переопределяем HtmlTemplatePath, если установлена переменная окружения
	htmlTemplatePath := os.Getenv("HTML_TEMPLATE_PATH")
	if htmlTemplatePath != "" {
		config.HtmlTemplatePath = htmlTemplatePath
	}

	return config, nil
}
