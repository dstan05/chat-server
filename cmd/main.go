package main

import (
	"github.com/dstan05/chat-server/internal/server"
)

func main() {
	s, err := server.Init()
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
}
