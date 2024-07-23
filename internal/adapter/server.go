package adapter

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GRPCAdapter struct {
	server   *grpc.Server
	grpcPort int
}

func NewGRPCAdapter(grpcPort int) *GRPCAdapter {
	return &GRPCAdapter{
		grpcPort: grpcPort,
	}
}

func (a *GRPCAdapter) Run() {
	log.Println("Starting gRPC server on port", a.grpcPort)
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", a.grpcPort))
	if err != nil {
		log.Fatalf("Error starting gRPC server: %v", err)
	}

	s := grpc.NewServer()
	a.server = s

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Error starting gRPC server: %v", err)
	}
}

func (a *GRPCAdapter) Stop() {
	a.server.Stop()
}
