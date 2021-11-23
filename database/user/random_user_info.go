package user

import "database/sql"

type RandomWallet struct {
	UserId int
	Wallet string
}

func GetRandomWallet(db *sql.DB) (RandomWallet, error) {
	const query = `SELECT id, wallet FROM blockchain.user ORDER BY RAND() LIMIT 1;`
	var randomWallet RandomWallet

	row, err := db.Query(query)
	if err != nil {
		return randomWallet, err
	}

	for row.Next() {
		var rw RandomWallet
		err = row.Scan(&rw.UserId, &rw.Wallet)
		if err != nil {
			return randomWallet, err
		}
		randomWallet = RandomWallet{rw.UserId, rw.Wallet}
	}

	return randomWallet, nil
}
