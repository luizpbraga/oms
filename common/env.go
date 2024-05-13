package common

import "syscall"

func Getenv(key, default_key string) string {
	if val, ok := syscall.Getenv(key); ok {
		return val
	}
	return default_key
}

