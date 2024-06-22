package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
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
	HtmlUiStaticPath string `json:"htmlUiStaticPath"`
}

func Read() (Config, error) {
	// Если выполняем "invest-tracker\cmd\web\ go run main.go",
	// то загрузятся переменные для локального окружения из .env
	// Если выполняем "docker-compose up --build",
	// то загрузятся переменные для контейнера из config.json

	// Загружаем переменные из файла .env
	envPath := "../../.env"
	if err := godotenv.Load(envPath); err != nil {
		fmt.Println("No .env file found")
	}

	var config Config
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
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

	overrideConfigFromEnv(&config)

	return config, nil
}

// Переопределяем config, если установлена переменная окружения
func overrideConfigFromEnv(config *Config) {
	if envVar := os.Getenv("HTML_TEMPLATE_PATH"); envVar != "" {
		config.HtmlTemplatePath = envVar
	}
	if envVar := os.Getenv("HTML_UISTATIC_PATH"); envVar != "" {
		config.HtmlUiStaticPath = envVar
	}
	if envVar := os.Getenv("DB_HOST"); envVar != "" {
		config.Database.Host = envVar
	}
	if envVar := os.Getenv("DB_PORT"); envVar != "" {
		if port, err := strconv.Atoi(envVar); err == nil {
			config.Database.Port = port
		}
	}
	if envVar := os.Getenv("DB_USER"); envVar != "" {
		config.Database.User = envVar
	}
	if envVar := os.Getenv("DB_PASS"); envVar != "" {
		config.Database.Password = envVar
	}
	if envVar := os.Getenv("DB_NAME"); envVar != "" {
		config.Database.DBName = envVar
	}
}
