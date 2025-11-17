package cache

import (
	"context"

	"github.com/go-redsync/redsync/v4"
)

type Client interface {
	LockClient
}

type LockClient interface {
	Lock(ctx context.Context, key string, opts ...redsync.Option) (LockClient, error)
	UnLock(ctx context.Context) error
}
