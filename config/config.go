package config

import (
	"github.com/spf13/cast"
	"os"
)

type Config struct {
	PostsServiceHost string

	HttpPort string

	Storage StorageConfig
}

type StorageConfig struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

func Load() Config {
	config := Config{}

	config.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8082"))
	config.PostsServiceHost = cast.ToString(getOrReturnDefault("POSTS_SERVICE_HOST", "localhost"))
	config.Storage.Host = cast.ToString(getOrReturnDefault("STORAGE_HOST", "localhost"))
	config.Storage.Port = cast.ToString(getOrReturnDefault("STORAGE_PORT", "5432"))
	config.Storage.Database = cast.ToString(getOrReturnDefault("STORAGE_DATABASE", "postgres"))
	config.Storage.Username = cast.ToString(getOrReturnDefault("STORAGE_USERNAME", "postgres"))
	config.Storage.Password = cast.ToString(getOrReturnDefault("STORAGE_PASSWORD", "postgres"))

	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
