package database

// func ConnectRedis() *redis.Client {
// 	REDIS_IP_PORT := os.Getenv("REDIS_IP_PORT")
// 	client := redis.NewClient(&redis.Options{
// 		Addr: REDIS_IP_PORT,
// 		DB:   0,
// 	})

// 	return client

// }

import (
	"context"
	"log"
	"os"
	"rediboard/models"

	"github.com/go-redis/redis/v8"
)

func NewDatabase(address string) (*models.Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return &models.Database{
		Client: client,
	}, nil
}

func RedisInstance() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_IP_PORT"),
		Password: "",
		DB:       0,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Failed to connect to redis: %s", err.Error())
		return nil
	}
	return client
}
