package migrations

import "database/sql"

type createBankExchangeRates struct{}

func getCreateBankExchange() migration {
	return &createBankExchangeRates{}
}

func (c *createBankExchangeRates) Name() string {
	return "create-bank-exchange-rates"
}

func (c *createBankExchangeRates) Up(conn *sql.Tx) error {
	_, err := conn.Exec(`CREATE TABLE bank_exchange_rates(
    	exchange_rate_uuid		UUID			PRIMARY KEY ,
    	from_currency			VARCHAR(5) 		NOT NULL ,
    	to_currency				VARCHAR(5) 		NOT NULL ,
    	rate					NUMERIC(20, 10)	NOT NULL ,
    	valid_from_timestamp	TIMESTAMP 		NOT NULL ,
    	valid_to_timestamp		TIMESTAMP 		NOT NULL ,
    	created_at 				TIMESTAMP 		NOT NULL DEFAULT NOW(),
		update_at 				TIMESTAMP 		NOT NULL DEFAULT NOW()
	)`)

	return err
}

func (c *createBankExchangeRates) Down(conn *sql.Tx) error {
	_, err := conn.Exec(`DROP TABLE bank_exchange_rates`)

	return err
}
