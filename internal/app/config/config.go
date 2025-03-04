package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/caarlos0/env/v6"
)

type AppConfig struct {
	ServiceURL  string `json:"service_url"`
	LogLevel    string `json:"log_level"`
	DatabaseURL string `json:"database_url"`
	RedisURL    string `json:"redis_url"`
}

func LoadConfig() (*AppConfig, error) {
	var configFile string

	cfg := &AppConfig{}

	// Обрабатываем флаги для выбора режима
	mode := flag.String("mode", "dev", "Set the mode (dev/prod)")
	flag.Parse()

	// Загружаем переменные окружения
	err := env.Parse(cfg)
	if err != nil {
		log.Fatal("Failed to parse environment variables:", err)
	}

	// Переопределяем конфиг, если есть переменные окружения
	if envServiceURL := os.Getenv("SERVER_URL"); envServiceURL != "" {
		cfg.ServiceURL = envServiceURL
	}

	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		cfg.LogLevel = envLogLevel
	}

	// Путь к конфигу зависит от режима
	if *mode == "dev" {
		configFile = "/Users/swarovski/Desktop/My life/Jobs/my-company/booking-service-go/internal/app/config/config-dev.json"
	} else if *mode == "prod" {
		configFile = "/Users/swarovski/Desktop/My life/Jobs/my-company/booking-service-go/internal/app/config/config-prod.json"
	} else {
		return nil, fmt.Errorf("invalid mode: %s", *mode)
	}

	// Чтение конфигурации из файла
	file, err := os.Open(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Декодируем JSON в структуру
	err = json.Unmarshal(content, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Обрабатываем флаг и переменные окружения
	if envAppMode := os.Getenv("APP_MODE"); envAppMode != "" {
		*mode = envAppMode
	}

	// Убедитесь, что значение конфигурации правильно прочитано из файла и переменных окружения
	log.Printf("Config loaded with mode: %s", *mode)

	return cfg, nil
}
