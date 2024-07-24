package adapter

import (
	"context"
	"github.com/c483481/bank_grpc_proto/protogen/go/bank"
	"google.golang.org/genproto/googleapis/type/date"
	"log"
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

func (a *GRPCAdapter) FetchExchangeRates(req *bank.ExchangeRateRequest, stream bank.BankService_FetchExchangeRatesServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("Client cancelled stream")
			return nil
		default:
			now := time.Now().Truncate(time.Second)
			rate := a.bankService.FindExchangeRate(req.FromCurrency, req.ToCurrency, now)

			err := stream.Send(&bank.ExchangeRateResponse{
				FromCurrency: req.FromCurrency,
				ToCurrency:   req.ToCurrency,
				Rate:         rate,
				Timestamp:    now.Format(time.RFC3339),
			})

			if err != nil {
				return err
			}

			log.Printf("Exchange rate send to client, %v to %v: %v", req.FromCurrency, req.ToCurrency, rate)

			time.Sleep(3 * time.Second)
		}
	}
}
