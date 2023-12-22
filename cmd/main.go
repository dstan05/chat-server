package main

import (
	"github.com/dstan05/chat-server/internal/config"
	"github.com/dstan05/chat-server/internal/pg"
	"github.com/dstan05/chat-server/internal/server"
)

func main() {
	c := &config.Config{}

	if err := c.Load(); err != nil {
		panic(err)
	}

	s, err := server.Init(&c.GrpcConfig)
	if err != nil {
		panic(err)
	}

	if err = s.Run(); err != nil {
		panic(err)
	}

	defer func() {
		if _, err := s.Stop(); err != nil {
			panic(err)
		}
	}()

	pgx, err := pg.Connect(c.PGConfig)
	defer func() {
		if err := pgx.Conn.Close(pgx.Ctx); err != nil {
			panic(err)
		}
	}()

}
