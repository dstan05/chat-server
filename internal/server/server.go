package server

import (
	"github.com/dstan05/chat-server/internal/config"
	"github.com/dstan05/chat-server/pkg/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Server struct {
	grps     *grpc.Server
	listener net.Listener
	config   config.GrpcConfig
}

func Init(c *config.GrpcConfig) (Server, error) {
	s := Server{
		config: *c,
	}
	s.grps = grpc.NewServer()

	reflection.Register(s.grps)
	chat.RegisterChatServer(s.grps, &Routes{})

	return s, nil
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", s.config.Address())
	if err != nil {
		return err
	}

	s.listener = lis
	if err = s.grps.Serve(lis); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() (Server, error) {
	s.grps.Stop()
	if err := s.listener.Close(); err != nil {
		return *s, err
	}

	return *s, nil
}
