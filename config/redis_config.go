package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func SetupRedisClient() {
	// Create a new Redis client
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"), // Replace with your Redis server address
		Password: "",                     // Replace with your Redis password, if any
		DB:       0,                      // Replace with your Redis database number
	})

	// Ping the Redis server to ensure the connection is successful
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pong, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		//return fmt.Errorf("failed to connect to Redis: %v", err)
		return
	}

	fmt.Println("Connected to Redis:", pong)

	return
}
