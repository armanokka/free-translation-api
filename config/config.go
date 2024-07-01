package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	/* Compulsory envs */
	RedisPassword string `yaml:"REDIS_PASSWORD" env:"REDIS_PASSWORD" env-required:"true"`
	RedisPort     int    `yaml:"REDIS_PORT" env:"REDIS_PORT" env-required:"true"`

	// Optional envs
	Port           int    `yaml:"PORT" env:"PORT" env-default:"80"`
	LogLevel       string `yaml:"LOG_LEVEL" env:"LOG_LEVEL" env-default:"debug"` // debug, info, error, fatal
	RedisDatabases int    `yaml:"REDIS_DATABASES" env:"REDIS_DATABASES" env-default:"1"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
