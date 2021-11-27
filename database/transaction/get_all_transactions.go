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

func GetLatestTransactions(db *sql.DB) (LatestTransactionsResponse, error) {
	const query = `SELECT time, amount FROM blockchain.transaction ORDER BY id DESC LIMIT 10;`
	var latestTransactionsResponse LatestTransactionsResponse

	row, err := db.Query(query)
	if err != nil {
		return latestTransactionsResponse, err
	}

	for row.Next() {
		var lt LatestTransactions
		err = row.Scan(&lt.Time, &lt.Amount)
		if err != nil {
			return latestTransactionsResponse, err
		}
		latestTransactionsResponse.Transactions = append(latestTransactionsResponse.Transactions,
			LatestTransactions{"", lt.Time, lt.Amount})
	}

	return latestTransactionsResponse, nil
}
