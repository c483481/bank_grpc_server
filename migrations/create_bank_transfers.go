package migrations

import "database/sql"

type createBankTransfers struct{}

func getCreateBankTransfers() migration {
	return &createBankTransfers{}
}

func (c *createBankTransfers) Name() string {
	return "create-bank-transfers"
}

func (c *createBankTransfers) Up(conn *sql.Tx) error {
	_, err := conn.Exec(`CREATE TABLE bank_transfers (
    	transfer_uuid	UUID 	NOT NULL PRIMARY KEY ,
    	from_account_uuid UUID 	NOT NULL REFERENCES bank_accounts (account_uuid),
    	to_account_uuid UUID NOT NULL  REFERENCES bank_accounts (account_uuid),
    	currency	VARCHAR(5) NOT NULL ,
    	amount		NUMERIC(15, 2) NOT NULL ,
    	transfer_success	BOOLEAN	NOT NULL DEFAULT FALSE,
    	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		update_at TIMESTAMP NOT NULL DEFAULT NOW()
	)`)

	return err
}

func (c *createBankTransfers) Down(conn *sql.Tx) error {
	_, err := conn.Exec(`DROP TABLE bank_transfers`)

	return err
}
