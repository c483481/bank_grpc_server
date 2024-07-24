package migrations

import "database/sql"

type createBankTransactions struct{}

func getCreateBankTransactions() migration {
	return &createBankTransactions{}
}

func (c *createBankTransactions) Name() string {
	return "create-bank-transactions"
}

func (c *createBankTransactions) Up(conn *sql.Tx) error {
	_, err := conn.Exec(`CREATE TABLE bank_transactions(
    	transaction_uuid	UUID	NOT NULL PRIMARY KEY,
    	account_uuid	UUID 	NOT NULL REFERENCES bank_accounts (account_uuid),
    	transaction_timestamp	TIMESTAMP NOT NULL,
    	amount	NUMERIC(15, 2)	NOT NULL ,
    	transactions_type VARCHAR(25) NOT NULL ,
    	notes	TEXT ,
    	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		update_at TIMESTAMP NOT NULL DEFAULT NOW()
	)`)
	return err
}

func (c *createBankTransactions) Down(conn *sql.Tx) error {
	_, err := conn.Exec(`DROP TABLE bank_transactions`)

	return err
}
