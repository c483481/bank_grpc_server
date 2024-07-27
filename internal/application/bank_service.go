package application

import (
	"fmt"
	"github.com/c483481/bank_grpc_server/database"
	"github.com/c483481/bank_grpc_server/internal/application/dto/bank"
	"github.com/c483481/bank_grpc_server/internal/types"
	"github.com/google/uuid"
	"log"
	"time"
)

type BankService struct {
	bankAccRepo     types.BankAccountDatabaseRepository
	exchangeRepo    types.BankExchangeRateDatabaseRepository
	transactionRepo types.TransactionDatabaseRepository
}

func GetBankService(bankRepo types.BankAccountDatabaseRepository, exchange types.BankExchangeRateDatabaseRepository, transaction types.TransactionDatabaseRepository) types.BankServiceType {
	return &BankService{
		bankAccRepo:     bankRepo,
		exchangeRepo:    exchange,
		transactionRepo: transaction,
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

func (s *BankService) CreateTransaction(acc string, t bank.Transaction) (uuid.UUID, error) {
	newUuid := uuid.New()
	now := time.Now()

	bankAccountOrm, err := s.bankAccRepo.GetBankAccountByAccountNumber(acc)

	if err != nil {
		log.Printf("Can't find account")
		return uuid.Nil, err
	}

	transactionOrm := database.BankTransactions{
		TransactionUuid:      newUuid,
		AccountUuid:          bankAccountOrm.AccountUuid,
		TransactionTimestamp: now,
		TransactionType:      t.TransactionType,
		Notes:                t.Notes,
		Amount:               t.Amount,
		UpdateAt:             now,
		CreatedAt:            now,
	}

	saveUuid, err := s.transactionRepo.CreateTransaction(bankAccountOrm, transactionOrm)

	return saveUuid, err
}

func (s *BankService) CalculateTransactionSummary(tcur *bank.TransactionSummary, trans bank.Transaction) error {
	switch trans.TransactionType {
	case bank.TransactionTypeIn:
		tcur.SumIn += trans.Amount
	case bank.TransactionTypeOut:
		tcur.SumOut += trans.Amount
	default:
		return fmt.Errorf("unkown transaction type %v", trans.TransactionType)
	}

	tcur.SumTotal = tcur.SumIn - tcur.SumOut

	return nil
}
