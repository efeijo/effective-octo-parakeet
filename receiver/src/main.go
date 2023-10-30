package main

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func main() {
	var ctx = context.Background()

	redisChan := os.Getenv("REDIS_CHANNEL")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	rdb := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: redisPassword,
		},
	)

	fmt.Println(rdb.Ping(ctx).Result())

	pubSub := rdb.Subscribe(ctx, redisChan)

	pubSub.Receive(ctx)
	for ev := range pubSub.Channel() {
		switch ev {

		}
	}
}
