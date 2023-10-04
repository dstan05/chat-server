package server

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	chat "github.com/dstan05/chat-server/pkg/grpc"
	"github.com/golang/protobuf/ptypes/empty"
)

type Routes struct {
	chat.UnimplementedChatServer
}

func (s *Routes) Create(ctx context.Context, req *chat.CreateRequest) (*chat.CreateResponse, error) {
	return &chat.CreateResponse{Id: gofakeit.Int64()}, nil
}

func (s *Routes) Delete(ctx context.Context, req *chat.DeleteRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *Routes) SendMessage(ctx context.Context, req *chat.SendMessageRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
