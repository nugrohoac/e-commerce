package constant

import "errors"

var (
	ErrInvalidLogin        = errors.New("invalid email/phone or password")
	ErrNothingToUnlock     = errors.New("nothing to unlock")
	ErrFailedToUnlockRedis = errors.New("failed to unlock redis")
	ErrNoWarehouseHasStock = errors.New("no warehouse has enough stock")
)
