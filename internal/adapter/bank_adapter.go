package adapter

import (
	"context"
	"github.com/c483481/bank_grpc_proto/protogen/go/bank"
	"google.golang.org/genproto/googleapis/type/date"
	"time"
)

func (a *GRPCAdapter) GetCurrentBalance(ctx context.Context, req *bank.CurrentBalanceRequest) (*bank.CurrentBalanceResponse, error) {
	now := time.Now()
	bal := a.bankService.FindCurrentBalance(req.AccountNumber)

	return &bank.CurrentBalanceResponse{
		Ammount: bal,
		CurrentDate: &date.Date{
			Year:  int32(now.Year()),
			Month: int32(now.Month()),
			Day:   int32(now.Day()),
		},
	}, nil
}
