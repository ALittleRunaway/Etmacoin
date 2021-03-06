package user

import (
	"database/sql"
)

func CheckUser(db *sql.DB, userPlain UserPlain) (User, error) {
	const query = `SELECT id, login, wallet, balance FROM blockchain.user WHERE login = ? and password = ?`
	var user User

	row, err := db.Query(query, userPlain.Login, userPlain.Password)
	if err != nil {
		return user, err
	}

	for row.Next() {
		var u User
		var up UserPlain
		err = row.Scan(&u.Id, &up.Login, &u.Wallet, &u.Balance)
		if err != nil {
			return user, err
		}
		user = User{UserPlain{u.Login, up.Password}, u.Id, u.Wallet, u.Balance}
	}
	return user, nil
}
