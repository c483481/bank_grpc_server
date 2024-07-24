package application

import (
	"github.com/c483481/bank_grpc_server/database"
	"github.com/c483481/bank_grpc_server/internal/application/dto/bank"
	"github.com/c483481/bank_grpc_server/internal/types"
	"github.com/google/uuid"
	"log"
	"time"
)

type BankService struct {
	bankAccRepo  types.BankAccountDatabaseRepository
	exchangeRepo types.BankExchangeRateDatabaseRepository
}

func GetBankService(bankRepo types.BankAccountDatabaseRepository, exchange types.BankExchangeRateDatabaseRepository) types.BankServiceType {
	return &BankService{
		bankAccRepo:  bankRepo,
		exchangeRepo: exchange,
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

func (s *BankService) CreateExchangeRate(r bank.ExchangeRate) (uuid.UUID, error) {
	newUuid := uuid.New()
	now := time.Now()

	exchange := database.BankExchangeRate{
		ExchangeRateUuid:   newUuid,
		FromCurrency:       r.FromCurrency,
		ToCurrency:         r.ToCurrency,
		Rate:               r.Rate,
		ValidFromTimestamp: r.ValidFromTimestamp,
		ValidToTimestamp:   r.ValidToTimestamp,
		UpdateAt:           now,
		CreatedAt:          now,
	}

	return s.exchangeRepo.CreateExchangeRate(exchange)
}

func (s *BankService) FindExchangeRate(from string, to string, ts time.Time) float64 {
	exchangeRate, err := s.exchangeRepo.GetExchangeRateAtTimestamp(from, to, ts)

	if err != nil {
		return 0
	}

	return exchangeRate.Rate
}
