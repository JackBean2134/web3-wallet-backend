package config

import (
	"os"
)

type Config struct {
	ServerPort string
	RPCURL     string
}

func LoadConfig() *Config {
	return &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		RPCURL:     getEnv("ETH_RPC_URL", "https://rpc.ankr.com/eth/0xE7381B9e9AaB14E90089Bba889D9e01fCa1F34bd"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
