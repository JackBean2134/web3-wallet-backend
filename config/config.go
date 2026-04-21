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
		RPCURL:     getEnv("ETH_RPC_URL", "https://rpc.ankr.com/eth/acth密钥"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
