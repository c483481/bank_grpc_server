package database

import (
	"github.com/google/uuid"
	"time"
)

type BankExchangeRate struct {
	ExchangeRateUuid   uuid.UUID `gorm:"column:exchange_rate_uuid;primaryKey;not null;<-:create"`
	FromCurrency       string    `gorm:"column:from_currency;not null"`
	ToCurrency         string    `gorm:"column:to_currency;not null"`
	Rate               float64   `gorm:"column:rate;not null"`
	ValidFromTimestamp time.Time `gorm:"column:valid_from_timestamp; not null"`
	ValidToTimestamp   time.Time `gorm:"column:valid_to_timestamp; not null"`
	UpdateAt           time.Time `gorm:"column:update_at;not null;autoCreateTime;autoUpdateTime"`
	CreatedAt          time.Time `gorm:"column:created_at;not null;autoCreateTime;<-:create"`
}

func (b *BankExchangeRate) TableName() string {
	return "bank_exchange_rates"
}
