package server

import (
	"fmt"
	"github.com/dstan05/chat-server/pkg/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

const port = 2020

type Server struct {
	grps     *grpc.Server
	listener net.Listener
}

func Init() (Server, error) {
	server := Server{}
	server.grps = grpc.NewServer()
	return server, nil
}

func (server *Server) Run() error {
	reflection.Register(server.grps)
	chat.RegisterChatServer(server.grps, &Routes{})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	server.listener = lis
	if err = server.grps.Serve(lis); err != nil {
		return err
	}

	return nil
}

func (server *Server) Stop() (Server, error) {
	server.grps.Stop()
	if err := server.listener.Close(); err != nil {
		return *server, err
	}

	return *server, nil
}
