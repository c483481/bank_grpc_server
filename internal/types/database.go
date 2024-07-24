package types

import "github.com/c483481/bank_grpc_server/database"

type BankAccountDatabaseRepository interface {
	GetBankAccountByAccountNumber(acc string) (*database.BankAccount, error)
}
