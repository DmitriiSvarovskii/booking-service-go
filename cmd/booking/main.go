package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/DmitriiSvarovskii/booking-service-go/internal/app/config"
	"github.com/DmitriiSvarovskii/booking-service-go/internal/app/logger"
	"github.com/DmitriiSvarovskii/booking-service-go/internal/app/server"
	"github.com/DmitriiSvarovskii/booking-service-go/internal/app/storage"
)

func main() {
	// Загружаем конфигурацию
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Инициализируем логирование
	if err := logger.Initialize(cfg.LogLevel); err != nil {
		log.Fatal(err)
	}

	// Подключаемся к PostgreSQL
	postgres := storage.NewPostgresClient(cfg.DatabaseURL)
	defer postgres.Close() // Обязательно закрываем соединение с БД

	// Подключаемся к Redis
	redis := storage.NewRedisClient(cfg.RedisURL)
	defer redis.Close() // Обязательно закрываем соединение с Redis

	// Настройка и запуск сервера
	srv := server.ShortenerRouter(cfg)
	go func() {
		if err := srv.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	// Обработка сигналов завершения работы
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Ожидаем сигнал завершения
	<-sigs
	log.Println("Shutting down gracefully...")
}
