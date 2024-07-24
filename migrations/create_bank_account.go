package migrations

import "database/sql"

type createBankAccount struct{}

func getCreateBankAcc() migration {
	return &createBankAccount{}
}

func (c *createBankAccount) Name() string {
	return "create-bank-account"
}

func (c *createBankAccount) Up(conn *sql.Tx) error {
	_, err := conn.Exec(`CREATE TABLE bank_accounts (
		account_uuid UUID NOT NULL PRIMARY KEY,
		account_number VARCHAR(20) UNIQUE NOT NULL,
		account_name VARCHAR(100) NOT NULL ,
		currency VARCHAR(5) NOT NULL ,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		update_at TIMESTAMP NOT NULL DEFAULT NOW()
	)`)
	return err
}

func (c *createBankAccount) Down(conn *sql.Tx) error {
	_, err := conn.Exec(`DROP TABLE bank_accounts`)

	return err
}
