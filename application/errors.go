package application

import "errors"

var (
	ErrLimitExceeded = errors.New("you have reached the maximum number of requests allowed within a certain time frame")
)
