package grpc

import (
	proto "github.com/RenterRus/sausage-profile/docs/proto/v1"
)

type Manager struct {
	proto.UnimplementedAuthServiceServer
}

func NewManager() proto.AuthServiceServer {
	return &Manager{}
}
