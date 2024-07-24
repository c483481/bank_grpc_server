package database

import (
	"github.com/google/uuid"
	"time"
)

type BankTransactions struct {
	TransactionUuid      uuid.UUID `gorm:"column:account_uuid;primaryKey;not null;<-:create"`
	AccountUuid          uuid.UUID `gorm:"column:account_uuid;unique;not null;<-:create"`
	TransactionTimestamp time.Time `gorm:"column:transaction_timestamp;not null;<-:create"`
	Amount               float64   `gorm:"column:amount;not null;<-:create"`
	TransactionType      string    `gorm:"column:transactions_type;not null;<-:create"`
	Notes                string    `gorm:"column:notes;not null;<-:create"`
	UpdateAt             time.Time `gorm:"column:update_at;not null;autoCreateTime;autoUpdateTime"`
	CreatedAt            time.Time `gorm:"column:created_at;not null;autoCreateTime;<-:create"`
}

func (t *BankTransactions) TableName() string {
	return "bank_transactions"
}
