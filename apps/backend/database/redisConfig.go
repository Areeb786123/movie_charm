package database

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         os.Getenv("REDIS_ADDR"),
		Password:     "",
		DB:           0,
		PoolSize:     20,
		MinIdleConns: 5,
	})
	ctx := context.Background()
	err := RedisClient.Ping(ctx).Err()
	if err != nil {
		log.Fatal("some error occur while connecting to redis", err)
	}
	log.Println("Redis connected successfully 💿 ")
}
