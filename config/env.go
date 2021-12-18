package config

import "os"

func env(key string, def string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return def
}
