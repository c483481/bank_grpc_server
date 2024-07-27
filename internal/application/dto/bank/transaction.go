package bank

import "time"

const (
	TransactionTypeUnknown string = "UNKNOWN"
	TransactionTypeIn      string = "IN"
	TransactionTypeOut     string = "OUT"
)

type Transaction struct {
	Amount          float64
	Timestamp       time.Time
	TransactionType string
	Notes           string
}

type TransactionSummary struct {
	SummaryOnDate time.Time
	SumIn         float64
	SumOut        float64
	SumTotal      float64
}
