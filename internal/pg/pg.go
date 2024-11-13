package pg

import (
	"context"
	"errors"
	"github.com/dstan05/chat-server/internal/config"
	"github.com/jackc/pgx/v5"
)

type PgxServer struct {
	Conn *pgx.Conn
	Ctx  context.Context
}

func Connect(config config.PGConfig) (PgxServer, error) {
	pg := PgxServer{
		Ctx: context.Background(),
	}

	err := errors.New("")
	pg.Conn, err = pgx.Connect(pg.Ctx, config.DSN())
	if err != nil {
		return pg, err
	}

	err = pg.Conn.Ping(pg.Ctx)
	if err != nil {
		return pg, err
	}

	return pg, nil
}
