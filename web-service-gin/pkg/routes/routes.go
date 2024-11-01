package routes

import (
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client // Global Redis client

func init() {
	// Initialize Redis client in init function
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
