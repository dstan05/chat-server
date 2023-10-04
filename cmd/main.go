package main

import (
	"github.com/dstan05/chat-server/internal/server"
)

func main() {
	s, err := server.Init()
	if err != nil {
		panic(err)
	}
	defer s.Stop()
}
