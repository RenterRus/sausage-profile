package grpc

import (
	"context"

	proto "github.com/RenterRus/sausage-profile/docs/proto/v1"
)

func (t *Manager) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	return &proto.RegisterResponse{
		Message: "unimplement",
	}, nil
}

func (t *Manager) Accept(ctx context.Context, req *proto.AcceptRequest) (*proto.AcceptResponse, error) {
	panic("unimplemented")
}
