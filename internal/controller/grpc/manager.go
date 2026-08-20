package grpc

import (
	proto "github.com/RenterRus/sausage-profile/docs/proto/v1"
	"github.com/RenterRus/sausage-profile/internal/usecase"
)

type Manager struct {
	proto.UnimplementedAuthServiceServer
	register usecase.Register
}

func NewManager(register usecase.Register) proto.AuthServiceServer {
	return &Manager{
		register: register,
	}
}
