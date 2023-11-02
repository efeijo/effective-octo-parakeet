package main

import (
	"context"
	"sender/internal/setup"
)

func main() {
	ctx := context.Background()
	rdb, err := setup.SetupRedis(ctx)

	if err != nil {
		panic(err)
	}

	rdb.Publish(ctx, "Success")

}
