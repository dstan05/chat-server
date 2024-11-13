package env

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

type PGConfig struct {
	dsn string
}

const (
	pgHost     = "PG_HOST"
	pgPort     = "PG_PORT"
	pgPassword = "PG_DB_NAME"
	pgUser     = "PG_USER"
	pgDbName   = "PG_PASSWORD"
)

func NewPGConfig(v *viper.Viper) (*PGConfig, error) {
	c := &PGConfig{}
	c.dsn = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		v.GetString(pgUser),
		v.GetString(pgPassword),
		v.GetString(pgHost),
		v.GetString(pgPort),
		v.GetString(pgDbName),
	)

	if len(c.dsn) == 0 {
		return c, errors.New("pg dsn not found")
	}

	return c, nil
}

func (cfg *PGConfig) DSN() string {
	return cfg.dsn
}
