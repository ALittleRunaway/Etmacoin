package transaction

import (
	"database/sql"
	"time"
)

type LatestTransactions struct {
	Hash   string
	Time   time.Time
	Amount int
}
type LatestTransactionsResponse struct {
	Transactions []LatestTransactions
}

type AllTransactionsResponse struct {
	Count        int
	Transactions []Transaction
}

func GetLatestTransactions(db *sql.DB) ([]Transaction, error) {
	const query = `SELECT id, sender_id, recipient_id, amount, message, time, prev_hash, pow
	FROM blockchain.transaction ORDER BY id DESC LIMIT 10;`
	var latestTransactions []Transaction

	row, err := db.Query(query)
	if err != nil {
		return latestTransactions, err
	}

	for row.Next() {
		var t Transaction
		err = row.Scan(&t.Id, &t.SenderId, &t.RecipientId, &t.Amount, &t.Message, &t.Time, &t.PrevHash, &t.PoW)
		if err != nil {
			return latestTransactions, err
		}
		var latestTransaction = Transaction{
			Id:          t.Id,
			SenderId:    t.SenderId,
			RecipientId: t.RecipientId,
			Amount:      t.Amount,
			Message:     t.Message,
			Time:        t.Time,
			PrevHash:    t.PrevHash,
			PoW:         t.PoW,
		}
		latestTransactions = append(latestTransactions, latestTransaction)
	}

	return latestTransactions, nil
}

func GetAllTransactions(db *sql.DB) ([]Transaction, error) {
	const query = `SELECT id, sender_id, recipient_id, amount, message, time, prev_hash, pow
	FROM blockchain.transaction WHERE id <> 1 ORDER BY id DESC;`
	var allTransactions []Transaction

	row, err := db.Query(query)
	if err != nil {
		return allTransactions, err
	}

	for row.Next() {
		var t Transaction
		err = row.Scan(&t.Id, &t.SenderId, &t.RecipientId, &t.Amount, &t.Message, &t.Time, &t.PrevHash, &t.PoW)
		if err != nil {
			return allTransactions, err
		}
		var transaction = Transaction{
			Id:          t.Id,
			SenderId:    t.SenderId,
			RecipientId: t.RecipientId,
			Amount:      t.Amount,
			Message:     t.Message,
			Time:        t.Time,
			PrevHash:    t.PrevHash,
			PoW:         t.PoW,
		}
		allTransactions = append(allTransactions, transaction)
	}

	return allTransactions, nil
}
