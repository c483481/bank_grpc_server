package types

import (
	"github.com/c483481/bank_grpc_server/database"
	"github.com/google/uuid"
	"time"
)

type BankAccountDatabaseRepository interface {
	GetBankAccountByAccountNumber(acc string) (*database.BankAccount, error)
}

type BankExchangeRateDatabaseRepository interface {
	CreateExchangeRate(r database.BankExchangeRate) (uuid.UUID, error)

	GetExchangeRateAtTimestamp(from string, to string, ts time.Time) (database.BankExchangeRate, error)
}

type TransactionDatabaseRepository interface {
	CreateTransaction(acc *database.BankAccount, t database.BankTransactions) (uuid.UUID, error)
}
