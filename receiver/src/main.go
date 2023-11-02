package main

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	redisHost := os.Getenv("MY_REDIS_MASTER_MASTER_SERVICE_HOST")
	redisChan := "channel"
	redisPassword := "password"

	fmt.Println("env vars", redisChan, redisPassword, redisHost)

	rdb := redis.NewClient(
		&redis.Options{
			Addr:     redisHost + ":6379",
			Password: redisPassword,
		},
	)

	_, err := rdb.Ping(ctx).Result()

	fmt.Println("ping", err)

	pubSub := rdb.Subscribe(ctx, redisChan)

	pubSub.Receive(ctx)
	for ev := range pubSub.Channel() {
		fmt.Println(ev.String())
	}
}
