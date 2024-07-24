package database

import (
	"github.com/google/uuid"
	"time"
)

type BankAccount struct {
	AccountUuid    uuid.UUID          `gorm:"column:account_uuid;primaryKey;not null;<-:create"`
	AccountNumber  string             `gorm:"column:account_number;unique;not null;<-:create"`
	AccountName    string             `gorm:"column:account_name;not null"`
	Currency       string             `gorm:"column:currency;not null"`
	CurrentBalance float64            `gorm:"column:current_balance;not null"`
	UpdateAt       time.Time          `gorm:"column:update_at;not null;autoCreateTime;autoUpdateTime"`
	CreatedAt      time.Time          `gorm:"column:created_at;not null;autoCreateTime;<-:create"`
	Transactions   []BankTransactions `gorm:"foreignKey:AccountUuid;references:AccountUuid"`
}

func (b *BankAccount) TableName() string {
	return "bank_accounts"
}
