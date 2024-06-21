package config

import "time"

type Config struct {
	LoggerConfig
}

type LoggerConfig struct {
	Level  string
	Format string
}

type HTTPClientConfig struct {
	Timeout time.Duration `env:"HTTP_CLIENT_TIMEOUT" envDefault:"60s"`
}

type DatabaseConfig struct {
	Path string `env:"DATABASE_PATH" envDefault:"./loft.db"`
}
