package setup

import (
	"context"
	"fmt"
	"sender/internal/sender"

	"os"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func SetupRedis(ctx context.Context, logger *zap.SugaredLogger) (sender.Sender, error) {

	redisHost := os.Getenv("MY_REDIS_MASTER_SERVICE_HOST")
	redisChan := "channel"
	redisPassword := "password"

	logger.Infoln("env vars", redisChan, redisPassword, redisHost)

	rdb := redis.NewClient(
		&redis.Options{
			Addr:     redisHost + ":6379",
			Password: redisPassword,
		},
	)

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		logger.Errorln("error pinging redis", err)
		return nil, fmt.Errorf("error pinging redis %w", err)
	}

	return sender.NewSender(rdb, redisChan), nil

}
