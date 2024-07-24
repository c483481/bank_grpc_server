package application

import (
	"github.com/c483481/bank_grpc_server/internal/types"
	"log"
)

type BankService struct {
	bankAccRepo types.BankAccountDatabaseRepository
}

func GetBankService(bankRepo types.BankAccountDatabaseRepository) types.BankServiceType {
	return &BankService{
		bankAccRepo: bankRepo,
	}
}

func (s *BankService) FindCurrentBalance(acc string) float64 {
	bankAccount, err := s.bankAccRepo.GetBankAccountByAccountNumber(acc)

	if err != nil {
		log.Println("Error FindCurrentBalance")
		return 0
	}

	return bankAccount.CurrentBalance
}
