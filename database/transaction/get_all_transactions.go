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

type UserTransaction struct {
	Direction    string
	CallerLogin  string
	CallerWallet string
	Amount       int
	Message      string
	Time         time.Time
}

type LatestTransactionsResponse struct {
	Transactions []LatestTransactions
}

type AllTransactionsResponse struct {
	Count        int
	Transactions []Transaction
}

type UserTransactionsResponse struct {
	Count        int
	Transactions []UserTransaction
}

func GetLatestTransactions(db *sql.DB) ([]Transaction, error) {
	const query = `SELECT id, COALESCE(sender_id, 0) as sender_id, COALESCE(recipient_id, 0) as recipient_id,
    amount, message, time, prev_hash, pow
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
	const query = `SELECT id, COALESCE(sender_id, 0) as sender_id, COALESCE(recipient_id, 0) as recipient_id, 
    amount, message, time, prev_hash, pow
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

func GetUserTransactions(db *sql.DB, userId int) ([]UserTransaction, error) {
	const query = `SELECT u.wallet, u.login, amount, message, time, 'me>' FROM blockchain.transaction t
		INNER JOIN blockchain.recipient r ON t.recipient_id = r.id
		INNER JOIN blockchain.sender s ON t.sender_id = s.id
		INNER JOIN blockchain.user u ON r.user_id = u.id
	WHERE t.sender_id = ?
	UNION
	SELECT u.wallet, u.login, amount, message, time, 'me<' FROM blockchain.transaction t
		INNER JOIN blockchain.recipient r ON t.recipient_id = r.id
		INNER JOIN blockchain.sender s ON t.sender_id = s.id
		INNER JOIN blockchain.user u ON s.user_id = u.id
	WHERE t.recipient_id = ?
	ORDER BY time DESC;`
	var userTransactions []UserTransaction

	row, err := db.Query(query, userId, userId)
	if err != nil {
		return userTransactions, err
	}

	for row.Next() {
		var ut UserTransaction
		err = row.Scan(&ut.CallerWallet, &ut.CallerLogin, &ut.Amount, &ut.Message, &ut.Time, &ut.Direction)
		if err != nil {
			return userTransactions, err
		}
		var transaction = UserTransaction{
			CallerWallet: ut.CallerWallet,
			CallerLogin:  ut.CallerLogin,
			Amount:       ut.Amount,
			Message:      ut.Message,
			Time:         ut.Time,
			Direction:    ut.Direction,
		}
		userTransactions = append(userTransactions, transaction)
	}

	return userTransactions, nil
}
