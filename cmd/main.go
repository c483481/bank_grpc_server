package main

import (
	"github.com/c483481/bank_grpc_server/internal/adapter"
)

func main() {
	grpcAdapter := adapter.NewGRPCAdapter(50000)

	grpcAdapter.Run()
}
