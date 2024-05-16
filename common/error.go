package common

import "errors"

var (
	ErrNoItems error = errors.New("Items len must be > 0")
)
