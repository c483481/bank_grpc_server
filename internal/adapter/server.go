package adapter

import (
	"fmt"
	"github.com/c483481/bank_grpc_proto/protogen/go/bank"
	"github.com/c483481/bank_grpc_server/internal/types"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GRPCAdapter struct {
	server      *grpc.Server
	bankService types.BankServiceType
	grpcPort    int
	bank.BankServiceServer
}

func NewGRPCAdapter(bankService types.BankServiceType, grpcPort int) *GRPCAdapter {
	return &GRPCAdapter{
		bankService: bankService,
		grpcPort:    grpcPort,
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

	bank.RegisterBankServiceServer(s, a)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Error starting gRPC server: %v", err)
	}
}

func (a *GRPCAdapter) Stop() {
	a.server.Stop()
}
