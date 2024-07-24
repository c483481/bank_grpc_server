package repository

import (
	"github.com/c483481/bank_grpc_server/database"
	"github.com/c483481/bank_grpc_server/internal/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BankExchangeRateRepository struct {
	db *gorm.DB
}

func GetExchangeRate(db *gorm.DB) types.BankExchangeRateDatabaseRepository {
	return &BankExchangeRateRepository{
		db: db,
	}
}

func (b BankExchangeRateRepository) CreateExchangeRate(r database.BankExchangeRate) (uuid.UUID, error) {
	if err := b.db.Create(r).Error; err != nil {
		return uuid.Nil, err
	}

	return r.ExchangeRateUuid, nil
}

func (b BankExchangeRateRepository) GetExchangeRateAtTimestamp(from string, to string, ts time.Time) (database.BankExchangeRate, error) {
	var result database.BankExchangeRate

	err := b.db.First(&result, "from_currency = ? AND to_currency = ? AND (? BETWEEN valid_from_timestamp and valid_to_timestamp)", from, to, ts).Error

	return result, err

}
