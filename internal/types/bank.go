package types

type BankServiceType interface {
	FindCurrentBalance(acc string) float64
}
