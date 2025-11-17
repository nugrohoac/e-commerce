package cache

import (
	"context"

	"github.com/go-redsync/redsync/v4"
	"github.com/nugrohoac/e-commerce/constant"
	"github.com/redis/go-redis/v9"
)

type client struct {
	db      *redis.Client
	redSync *redsync.Redsync
	mutex   *redsync.Mutex
}

func (c client) Lock(ctx context.Context, key string, opts ...redsync.Option) (LockClient, error) {
	mutex := c.redSync.NewMutex(key, opts...)
	if err := mutex.LockContext(ctx); err != nil {
		return nil, err
	}

	return &client{
		db:      c.db,
		redSync: c.redSync,
		mutex:   mutex,
	}, nil
}

func (c client) UnLock(ctx context.Context) error {
	if c.mutex == nil {
		return constant.ErrNothingToUnlock
	}

	ok, err := c.mutex.UnlockContext(ctx)
	if !ok || err != nil {
		if err == nil {
			err = constant.ErrFailedToUnlockRedis
		}
		return err
	}

	return nil
}

func NewClient(db *redis.Client, redSync *redsync.Redsync) Client {
	return client{
		db:      db,
		redSync: redSync,
	}
}
