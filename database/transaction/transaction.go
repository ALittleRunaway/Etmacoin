package database

import (
	"Blockchain/database"
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
	Id          int
	SenderId    int
	RecipientId int
	Amount      int
	Message     string
	Time        time.Time
	PrevHash    string
	PoW         int
}

type AddNewTransactionResponse struct {
	Transaction
	Response string
}

func AddNewTransactionHandler(newTransactionPlain TransactionPlain) (AddNewTransactionResponse, error) {
	db, err := database.Connection()
	var newTransaction Transaction
	var newTransactionResponse AddNewTransactionResponse
	newTransaction.Amount = newTransactionPlain.Amount
	newTransaction.Message = newTransactionPlain.Message
	if err != nil {
		return newTransactionResponse, err
	}
	newTransaction.SenderId, err = GetSenderId(db, newTransactionPlain.UserId)
	if err != nil {
		return newTransactionResponse, err
	}
	newTransaction.RecipientId, err = GetRecipientId(db, newTransactionPlain.RecipientWallet)
	if err != nil {
		return newTransactionResponse, err
	}
	newTransaction.PrevHash = "gjkhjk"
	newTransaction.PoW = 0
	newTransaction.Time = timeIn("Europe/Moscow")
	err = AddNewTransaction(db, newTransaction)
	if err != nil {
		return newTransactionResponse, err
	}
	return newTransactionResponse, nil
}

func GetSenderId(db *sql.DB, userId int) (int, error) {
	const query = `SELECT id FROM blockchain.sender s WHERE s.user_id = ?`
	senderId := 0

	row, err := db.Query(query, userId)
	if err != nil {
		return senderId, err
	}

	for row.Next() {
		var i int
		err = row.Scan(&i)
		if err != nil {
			return senderId, err
		}
		senderId = i
	}

	return senderId, nil
}

func GetRecipientId(db *sql.DB, recipientWallet string) (int, error) {
	const query = `SELECT id FROM blockchain.recipient r WHERE r.user_id = (
    	SELECT id FROM blockchain.user WHERE wallet = ?)`
	RecipientId := 0

	row, err := db.Query(query, recipientWallet)
	if err != nil {
		return RecipientId, err
	}

	for row.Next() {
		var i int
		err = row.Scan(&i)
		if err != nil {
			return RecipientId, err
		}
		RecipientId = i
	}

	return RecipientId, nil
}

func AddNewTransaction(db *sql.DB, newTransaction Transaction) error {
	db, err := database.Connection()
	if err != nil {
		return err
	}

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

func timeIn(name string) time.Time {
	loc, err := time.LoadLocation(name)
	if err != nil {
		panic(err)
	}
	return time.Now().In(loc)
}
