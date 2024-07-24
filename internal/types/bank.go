package types

import (
	"github.com/c483481/bank_grpc_server/internal/application/dto/bank"
	"github.com/google/uuid"
	"time"
)

type BankServiceType interface {
	FindCurrentBalance(acc string) float64

	CreateExchangeRate(r bank.ExchangeRate) (uuid.UUID, error)

	FindExchangeRate(from string, to string, ts time.Time) float64
}
