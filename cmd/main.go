package main

import (
	"github.com/c483481/bank_grpc_server/internal/adapter"
	"github.com/c483481/bank_grpc_server/internal/application"
)

func main() {
	bs := &application.BankService{}
	grpcAdapter := adapter.NewGRPCAdapter(bs, 50000)

	grpcAdapter.Run()
}
