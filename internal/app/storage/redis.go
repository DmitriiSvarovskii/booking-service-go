package storage

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(redisURL string) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr: redisURL,
	})

	// Проверка соединения с Redis
	_, err := client.Ping(context.Background()).Result()

	if err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}
	log.Println("Successfully connected to redis")

	return &RedisClient{client: client}
}

func (r *RedisClient) Close() {
	err := r.client.Close()
	if err != nil {
		log.Printf("Error closing Redis connection: %v", err)
	}
}
