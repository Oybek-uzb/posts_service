package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	PostsServiceHost string
	PostsServicePort int

	HttpPort string

	Storage StorageConfig `json:"storage"`
}

type StorageConfig struct {
	Host     string `json:"host" env-default:"localhost"`
	Port     string `json:"port" env-default:"5432"`
	Database string `json:"database" env-default:"posts-service"`
	Username string `json:"username" env-default:"postgres"`
	Password string `json:"password" env-default:"postgres"`
}

func Load() Config {
	config := Config{}

	config.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8082"))

	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
