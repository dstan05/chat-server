package server

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/dstan05/chat-server/pkg/chat"
	"github.com/golang/protobuf/ptypes/empty"
)

type Routes struct {
	chat.UnimplementedChatServer
}

func (r *Routes) Create(_ context.Context, _ *chat.CreateRequest) (*chat.CreateResponse, error) {
	return &chat.CreateResponse{Id: gofakeit.Int64()}, nil
}

func (r *Routes) Delete(_ context.Context, _ *chat.DeleteRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (r *Routes) SendMessage(_ context.Context, _ *chat.SendMessageRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
