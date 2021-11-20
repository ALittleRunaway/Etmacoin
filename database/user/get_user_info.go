package database

import (
	"Blockchain/database"
	"database/sql"
)

type UserInfo struct {
	Login   string
	Wallet  string
	Balance int
}

func GetUserInfoHandler(userId int) (UserInfo, error) {
	db, err := database.Connection()
	var userInfo UserInfo
	if err != nil {
		return userInfo, err
	}
	userInfo, err = GetUserInfo(db, userId)
	if err != nil {
		return userInfo, err
	}
	return userInfo, nil
}

func GetUserInfo(db *sql.DB, userId int) (UserInfo, error) {
	const query = `SELECT login, wallet, balance FROM blockchain.user WHERE id = ?`
	var userInfo UserInfo

	row, err := db.Query(query, userId)
	if err != nil {
		return userInfo, err
	}

	for row.Next() {
		var ui UserInfo
		err = row.Scan(&ui.Login, &ui.Wallet, &ui.Balance)
		if err != nil {
			return userInfo, err
		}
		userInfo = UserInfo{ui.Login, ui.Wallet, ui.Balance}
	}

	return userInfo, nil
}
