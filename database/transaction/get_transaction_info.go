package transaction

import (
	"database/sql"
)

func GetLastTransaction(db *sql.DB) (Transaction, error) {
	const query = `SELECT id, sender_id, recipient_id, amount, message, time, prev_hash, pow
	FROM blockchain.transaction ORDER BY ID DESC LIMIT 1`
	var lastTransaction Transaction

	row, err := db.Query(query)
	if err != nil {
		return lastTransaction, err
	}

	for row.Next() {
		var lt Transaction
		err = row.Scan(&lt.Id, &lt.SenderId, &lt.RecipientId, &lt.Amount, &lt.Message, &lt.Time, &lt.PrevHash, &lt.PoW)
		if err != nil {
			return lastTransaction, err
		}
		lastTransaction = Transaction{
			Id:              lt.Id,
			SenderId:        lt.SenderId,
			SenderUserId:    0,
			RecipientId:     lt.Amount,
			RecipientUserId: 0,
			Amount:          lt.Amount,
			Message:         lt.Message,
			Time:            lt.Time,
			PrevHash:        lt.PrevHash,
			PoW:             lt.PoW,
		}
	}

	return lastTransaction, nil
}
