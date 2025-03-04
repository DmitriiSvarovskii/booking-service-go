package storage

import (
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PostgresClient struct {
	db *gorm.DB
}

func NewPostgresClient(dsn string) *PostgresClient {
	// Логируем строку подключения
	log.Printf("Connecting to database with DSN: %s", dsn)
	// Добавляем sslmode=disable
	if !strings.Contains(dsn, "sslmode") {
		dsn += "?sslmode=disable"
	}
	// Открываем соединение с базой данных
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	log.Println("Successfully connected to database")
	return &PostgresClient{db: db}
}

func (p *PostgresClient) Close() {
	err := p.db.Close()
	if err != nil {
		log.Printf("Error closing PostgreSQL connection: %v", err)
	}
}
