package sender

import (
	"context"
	"encoding/json"
	"sender/internal/model"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

const (
	SET = "sender-set"
)

type conf struct {
	redisChan string
}

type sender struct {
	rdb  *redis.Client
	conf *conf
}

type Sender interface {
	Publish(ctx context.Context, message *model.Message, logger *zap.SugaredLogger) error
	GetAllPublished(ctx context.Context, logger *zap.SugaredLogger) ([]model.Message, error)
}

func NewSender(client *redis.Client, redisChan string) Sender {
	return &sender{
		rdb: client,
		conf: &conf{
			redisChan: redisChan,
		},
	}
}

func (s *sender) Publish(ctx context.Context, message *model.Message, logger *zap.SugaredLogger) error {
	l := logger
	err := s.rdb.Publish(ctx, s.conf.redisChan, message).Err()
	if err != nil {
		l.Errorln("error publishing", err)
		return err
	}

	err = s.rdb.SAdd(ctx, SET, message).Err()
	if err != nil {
		l.Errorln("error adding to set", err)
		return err
	}
	return nil
}

// GetAllPublished implements Sender.
func (s *sender) GetAllPublished(ctx context.Context, logger *zap.SugaredLogger) ([]model.Message, error) {
	l := logger
	res, err := s.rdb.SMembers(ctx, SET).Result()
	if err != nil {
		l.Errorln("error getting messages", err)
		return nil, err
	}
	messages := make([]model.Message, 0, len(res))
	for _, mes := range res {
		var message model.Message
		err = json.Unmarshal([]byte(mes), &message)
		if err != nil {
			l.Errorln("error unmarshaling message: "+mes, err)
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}
