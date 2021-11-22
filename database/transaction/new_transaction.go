package database

import (
	"database/sql"
	"time"
)

type TransactionPlain struct {
	UserId          int
	RecipientWallet string
	Amount          int
	Message         string
}

type Transaction struct {
	Id              int
	SenderId        int
	SenderUserId    int
	RecipientId     int
	RecipientUserId int
	Amount          int
	Message         string
	Time            time.Time
	PrevHash        string
	PoW             int
}

type AddNewTransactionResponse struct {
	Transaction
	Response string
}

func AddNewTransaction(db *sql.DB, newTransaction Transaction) error {

	const query = `INSERT INTO blockchain.transaction 
    (sender_id, recipient_id, amount, message, time, prev_hash, pow) VALUES (?, ?, ?, ?, ?, ?, ?)`

	queryConn, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer queryConn.Close()

	if _, err = queryConn.Exec(newTransaction.SenderId, newTransaction.RecipientId, newTransaction.Amount,
		newTransaction.Message, newTransaction.Time, newTransaction.PrevHash, newTransaction.PoW); err != nil {
		return err
	}
	return nil
}

func TakeCoinsFromSender(db *sql.DB, newTransaction Transaction) error {

	// Get current sender balance
	var query1 = `SELECT u.balance FROM blockchain.user u WHERE u.id = ?`

	oldBalance := 0

	row, err := db.Query(query1, newTransaction.SenderUserId)
	if err != nil {
		return err
	}

	for row.Next() {
		var i int
		err = row.Scan(&i)
		if err != nil {
			return err
		}
		oldBalance = i
	}

	// Take coins from sender balance
	var query2 = `UPDATE blockchain.user u SET u.balance = ? - ? WHERE u.id = ?`

	queryConn, err := db.Prepare(query2)
	if err != nil {
		return err
	}
	defer queryConn.Close()

	if _, err = queryConn.Exec(oldBalance, newTransaction.Amount, newTransaction.SenderUserId); err != nil {
		return err
	}
	return nil
}

func AddCoinsToRecipient(db *sql.DB, newTransaction Transaction) error {

	// Get current recipient balance
	var query1 = `SELECT u.balance FROM blockchain.user u WHERE u.id = ?`

	oldBalance := 0

	row, err := db.Query(query1, newTransaction.RecipientUserId)
	if err != nil {
		return err
	}

	for row.Next() {
		var i int
		err = row.Scan(&i)
		if err != nil {
			return err
		}
		oldBalance = i
	}

	// Add coins to the recipient's balance
	var query2 = `UPDATE blockchain.user u SET u.balance = ? + ? WHERE u.id = ?`

	queryConn, err := db.Prepare(query2)
	if err != nil {
		return err
	}
	defer queryConn.Close()

	if _, err = queryConn.Exec(oldBalance, newTransaction.Amount, newTransaction.RecipientUserId); err != nil {
		return err
	}
	return nil
}
