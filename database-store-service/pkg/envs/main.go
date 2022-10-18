package envs

import "os"

func Getenv(value string, fallback string) string {
	if value, ok := os.LookupEnv(value); ok {
		return value
	}
	return fallback
}