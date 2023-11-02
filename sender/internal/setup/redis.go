package setup

import (
	"context"
	"fmt"
	"os"
	"sender/internal/sender"

	"github.com/redis/go-redis/v9"
)

func SetupRedis(ctx context.Context) (*sender.Sender, error) {

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

	if err != nil {
		return nil, fmt.Errorf("error pinging redis %w", err)
	}

	return sender.NewSender(rdb, redisChan), nil

}
