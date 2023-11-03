package main

import (
	"context"
	"sender/internal/router"
	"sender/internal/setup"

	"net/http"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	REDIS_CHANNEL = "channel"
	PORT          = ":8080"
)

func main() {
	logger, _ := zap.NewDevelopment(
		zap.WithClock(zapcore.DefaultClock),
	)
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	ctx := context.Background()

	send, err := setup.SetupRedis(ctx, sugar)
	if err != nil {
		sugar.DPanicln("error initializing redis", err)
	}

	r := router.NewRouter(send, sugar)

	sugar.Infoln("Listening on port: ", PORT)
	http.ListenAndServe(PORT, r)
}
