package main

import (
	"flag"
	"github.com/c483481/bank_grpc_server/database"
	"github.com/c483481/bank_grpc_server/internal/adapter"
	"github.com/c483481/bank_grpc_server/internal/application"
	"github.com/c483481/bank_grpc_server/internal/application/dto/bank"
	"github.com/c483481/bank_grpc_server/internal/repository"
	"github.com/c483481/bank_grpc_server/internal/types"
	"github.com/c483481/bank_grpc_server/migrations"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	log.Println("Loading environment variables...")
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file, attempting to use environment variables")
	}

	sqlUri := os.Getenv("POSTGRES_URI")

	if sqlUri == "" {
		log.Fatal("Error get POSTGRES_URI from env.")
	}

	db := database.GetDatabase(sqlUri)

	sqlDb, _ := db.DB()

	downFlag := flag.Bool("down", false, "Run database migration down")
	downAllFlag := flag.Bool("down-all", false, "Run all database migrations down")

	// Parse the flags
	flag.Parse()

	if *downFlag {
		log.Println("Running database migration down...")
		migrations.Down(sqlDb)
		log.Println("Successfully run database migration down.")
		return
	}

	if *downAllFlag {
		log.Println("Running all database migrations down...")
		migrations.DownAll(sqlDb)
		log.Println("Successfully run all database migrations down.")
		return
	}

	log.Println("Running database migration up...")

	migrations.Up(sqlDb)

	log.Println("Successfully run database migration up.")

	rba := repository.GetBankAccountRepository(db)
	eba := repository.GetExchangeRate(db)
	tb := repository.NewBankTransaction(db)

	bs := application.GetBankService(rba, eba, tb)

	//go generateExchanegRate(bs, "USD", "IDR", 5*time.Second)

	grpcAdapter := adapter.NewGRPCAdapter(bs, 50000)

	grpcAdapter.Run()
}

func generateExchanegRate(bs types.BankServiceType, from, to string, duration time.Duration) {
	ticker := time.NewTicker(duration)

	for range ticker.C {
		now := time.Now()
		validFrom := now.Truncate(time.Second).Add(3 * time.Second)
		validTo := validFrom.Add(duration).Add(-1 * time.Millisecond)

		dumyRate := bank.ExchangeRate{
			FromCurrency:       from,
			ToCurrency:         to,
			ValidFromTimestamp: validFrom,
			ValidToTimestamp:   validTo,
			Rate:               2000 + float64(rand.Intn(300)),
		}

		uuid, err := bs.CreateExchangeRate(dumyRate)
		if err != nil {
			log.Println("failed to insert exchange rate. error : ", err)
			continue
		}
		log.Println("success add exchange rate, uuid : ", uuid)
	}
}
