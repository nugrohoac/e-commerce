package constant

import "errors"

var (
	ErrInvalidLogin = errors.New("invalid email/phone or password")
)
