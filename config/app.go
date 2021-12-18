package config

var (
	Port = env("HISTORY_SERVER_LISTEN_ADDR", "8080")
)