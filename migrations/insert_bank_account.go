package migrations

import "database/sql"

type insertBankAccount struct{}

func getInsertBankAccount() migration {
	return &insertBankAccount{}
}

func (i *insertBankAccount) Name() string {
	return "insert-bank-account"
}

func (i *insertBankAccount) Up(conn *sql.Tx) error {
	_, err := conn.Exec(`
		INSERT INTO bank_accounts VALUES 
		('e92c0c1e-6ff9-449e-90af-daf9beecda7c', '5432101', 'John Doe', 'USD', 10, NOW(), NOW()),
		('17b28816-6b1e-400c-90b8-f2f111f42299', '5432102', 'Jane Doe', 'USD', 10, NOW(), NOW()),
		('0e0a4f44-9b24-48fd-a617-fc0579e6d098', '5432103', 'Bob Smith', 'USD', 10, NOW(), NOW()),
		('ea86e70a-7529-4c48-b7e3-49b8968d587e', '5432104', 'Alice Johnson', 'USD', 10, NOW(), NOW()),
		('3f246106-33bc-419c-812e-4164a98acae0', '5432105', 'Emily Davis', 'USD', 10, NOW(), NOW())
	`)

	return err
}

func (i *insertBankAccount) Down(conn *sql.Tx) error {
	_, err := conn.Exec(`DELETE FROM bank_accounts WHERE account_uuid IN 
    	('e92c0c1e-6ff9-449e-90af-daf9beecda7c', '17b28816-6b1e-400c-90b8-f2f111f42299', '0e0a4f44-9b24-48fd-a617-fc0579e6d098', 'ea86e70a-7529-4c48-b7e3-49b8968d587e', '3f246106-33bc-419c-812e-4164a98acae0')
	`)

	return err
}
