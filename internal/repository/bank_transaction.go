package repository

import (
	"github.com/c483481/bank_grpc_server/database"
	"github.com/c483481/bank_grpc_server/internal/application/dto/bank"
	"github.com/c483481/bank_grpc_server/internal/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BankDatabaseTransaction struct {
	db *gorm.DB
}

func NewBankTransaction(db *gorm.DB) types.TransactionDatabaseRepository {
	return &BankDatabaseTransaction{
		db: db,
	}
}

func (b *BankDatabaseTransaction) CreateTransaction(acc *database.BankAccount, t database.BankTransactions) (uuid.UUID, error) {
	tx := b.db.Begin()

	if err := tx.Create(t).Error; err != nil {
		tx.Rollback()
		return uuid.Nil, err
	}

	newAmount := t.Amount
	if t.TransactionType == bank.TransactionTypeOut {
		newAmount = -1 * t.Amount
	}

	newAccountBalance := acc.CurrentBalance + newAmount

	if err := tx.Model(acc).Updates(map[string]interface{}{
		"current_balance": newAccountBalance,
		"update_at":       time.Now(),
	}).Error; err != nil {
		tx.Rollback()
		return uuid.Nil, err
	}

	tx.Commit()

	return t.TransactionUuid, nil
}
