package app

import (
	"fmt"
	"log"
	"net"

	"github.com/sourcegraph/conc/pool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	grpcAddr string
}

func NewApp(configPath string) (*App, error) {
	lastSlash := 0
	for i, v := range configPath {
		if v == '/' {
			lastSlash = i
		}
	}

	conf, err := ReadConfig(configPath[:lastSlash], configPath[lastSlash+1:])
	if err != nil {
		return nil, fmt.Errorf("ReadConfig: %w", err)
	}

	_ = conf

	return &App{
		grpcAddr: fmt.Sprintf("%s:%d", conf.GRPC.Host, conf.GRPC.Port),
	}, nil
}

func (a *App) Run() error {
	p := pool.New().WithErrors()

	// grpc
	p.Go(func() error {
		lis, err := net.Listen("tcp", a.grpcAddr)
		if err != nil {
			return fmt.Errorf("failed to listen: %v", err)
		}

		s := grpc.NewServer()
		reflection.Register(s)

		defer func() {
			s.Stop()
		}()

		//v1.RegisterAuthServiceServer(s, protoServe.NewManager())

		log.Printf("gRPC server listening on %s", a.grpcAddr)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

		return nil
	})

	if err := p.Wait(); err != nil {
		return fmt.Errorf("Run: %w", err)
	}

	return nil
}
