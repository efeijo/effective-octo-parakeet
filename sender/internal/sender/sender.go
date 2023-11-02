package sender

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type conf struct {
	redisChan string
}

type Sender struct {
	rdb  *redis.Client
	conf *conf
}

func NewSender(client *redis.Client, redisChan string) *Sender {
	return &Sender{
		rdb: client,
		conf: &conf{
			redisChan: redisChan,
		},
	}
}

func (s *Sender) Publish(ctx context.Context, message any) error {
	return s.rdb.Publish(ctx, s.conf.redisChan, time.Now().String()).Err()
}
