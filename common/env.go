package common

import "syscall"

// EnvString retrieves the value of the environment variable named by the key.
func EnvString(key, fallback string) string {
	// Use syscall.Getenv to get the environment variable value
	if val, ok := syscall.Getenv(key); ok {
		return val
	}
	// If the environment variable is not set, return the fallback value, fallback value is used when the environment variable is not set, it can be a default value or an empty string
	return fallback
}
