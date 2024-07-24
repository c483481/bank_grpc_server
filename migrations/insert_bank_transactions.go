package migrations

import "database/sql"

type insertBankTransactions struct{}

func getInsertBankTransactions() migration {
	return &insertBankTransactions{}
}

func (i *insertBankTransactions) Name() string {
	return "insert-bank-transactions"
}

func (i *insertBankTransactions) Up(conn *sql.Tx) error {
	_, err := conn.Exec(`INSERT INTO bank_transactions VALUES 
		('c524c7a9-4352-40a5-9d0e-35251a206310', 'e92c0c1e-6ff9-449e-90af-daf9beecda7c', NOW(), 10, 'IN', 'Initial Deposit', NOW(), NOW()),
		('9783e9aa-333e-44b6-ab58-266d19d609b6', '17b28816-6b1e-400c-90b8-f2f111f42299', NOW(), 10, 'IN', 'Initial Deposit', NOW(), NOW()),
		('b81281e5-17ca-4fe6-befa-124c891c1c60', '0e0a4f44-9b24-48fd-a617-fc0579e6d098', NOW(), 10, 'IN', 'Initial Deposit', NOW(), NOW()),
		('0f74895e-ee25-4838-af10-8a7df4aafaca', 'ea86e70a-7529-4c48-b7e3-49b8968d587e', NOW(), 10, 'IN', 'Initial Deposit', NOW(), NOW()),
		('89f3f060-b315-487f-8924-d90417471cab', '3f246106-33bc-419c-812e-4164a98acae0', NOW(), 10, 'IN', 'Initial Deposit', NOW(), NOW())
	`)

	return err
}

func (i *insertBankTransactions) Down(conn *sql.Tx) error {
	_, err := conn.Exec(`DELETE FROM bank_transactions WHERE transaction_uuid IN 
		('c524c7a9-4352-40a5-9d0e-35251a206310', '9783e9aa-333e-44b6-ab58-266d19d609b6', 'b81281e5-17ca-4fe6-befa-124c891c1c60', '0f74895e-ee25-4838-af10-8a7df4aafaca', '89f3f060-b315-487f-8924-d90417471cab')	
	`)

	return err
}
