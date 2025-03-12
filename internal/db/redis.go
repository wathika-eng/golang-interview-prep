package db

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

// initializes a redis instance
func RedisInit() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     env.REDIS_URL, // Redis server address
		Password: "",            // No password set
		DB:       0,             // Use default DB
	})

	// Ping the Redis server to check the connection
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v\n", err)
	}

	fmt.Println("Successfully connected to Redis!")
	return client
}
