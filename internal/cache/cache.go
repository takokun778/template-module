package cache

import (
	"context"
	"fmt"

	"github.com/redis/rueidis"
)

type Cache struct {
	Client rueidis.Client
}

func New(
	url string,
) (*Cache, error) {
	ctx := context.Background()

	option := rueidis.ClientOption{
		InitAddress: []string{url},
	}

	cli, err := rueidis.NewClient(option)
	if err != nil {
		return nil, fmt.Errorf("failed to new redis: %w", err)
	}

	if err := cli.Do(ctx, cli.B().Ping().Build()).Error(); err != nil {
		return nil, fmt.Errorf("failed to new redis: %w", err)
	}

	return &Cache{
		Client: cli,
	}, nil
}
