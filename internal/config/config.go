package config

import (
	"github.com/dstan05/chat-server/internal/config/env"
	"github.com/spf13/viper"
)

type GrpcConfig interface {
	Address() string
}

type PGConfig interface {
	DSN() string
}

type Config struct {
	ConfigFile string
	GrpcConfig GrpcConfig
	PGConfig   PGConfig
}

func (c *Config) Load() error {
	if c.ConfigFile == "" {
		c.ConfigFile = "./.env"
	}

	v := viper.New()
	v.SetConfigFile(c.ConfigFile)

	err := v.ReadInConfig()
	if err != nil {
		return err
	}

	gc, err := env.NewGRPCConfig(v)
	if err != nil {
		return err
	}

	pgc, err := env.NewPGConfig(v)
	if err != nil {
		return err
	}

	c.GrpcConfig = gc
	c.PGConfig = pgc

	return nil
}

var _ GrpcConfig = (*env.GrpcConfig)(nil)

var _ PGConfig = (*env.PGConfig)(nil)
