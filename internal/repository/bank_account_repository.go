package repository

import (
	"github.com/c483481/bank_grpc_server/database"
	"github.com/c483481/bank_grpc_server/internal/types"
	"gorm.io/gorm"
	"log"
)

type BankAccountRepository struct {
	db *gorm.DB
}

func GetBankAccountRepository(db *gorm.DB) types.BankAccountDatabaseRepository {
	return &BankAccountRepository{
		db: db,
	}
}

func (b *BankAccountRepository) GetBankAccountByAccountNumber(acc string) (*database.BankAccount, error) {
	var result database.BankAccount

	if err := b.db.First(&result, "account_number = ?", acc).Error; err != nil {
		log.Printf("Can't find bank account number %s\n", acc)
		return nil, err
	}

	return &result, nil
}
