package database

import (
	"context"
	"os"

	"github.com/go-redis/redis"
)

var Ctx = context.Background()

func CreateClient(dbNum int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DATABASE_ADDRESS"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		DB:       dbNum,
	})

	return rdb
}
