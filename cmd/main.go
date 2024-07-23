package main

import (
	"flag"
	"github.com/c483481/bank_grpc_server/database"
	"github.com/c483481/bank_grpc_server/internal/adapter"
	"github.com/c483481/bank_grpc_server/internal/application"
	"github.com/c483481/bank_grpc_server/migrations"
	"github.com/joho/godotenv"
	"log"
	"os"
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

	bs := &application.BankService{}
	grpcAdapter := adapter.NewGRPCAdapter(bs, 50000)

	grpcAdapter.Run()
}
