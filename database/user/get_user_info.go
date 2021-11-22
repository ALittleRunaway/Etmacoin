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

func GetRecipientAndUserId(db *sql.DB, recipientWallet string) (int, int, error) {
	const query = `SELECT r.id, r.user_id FROM blockchain.recipient r WHERE r.user_id = (
    	SELECT id FROM blockchain.user WHERE wallet = ?)`
	RecipientId := 0
	RecipientUserId := 0

	row, err := db.Query(query, recipientWallet)
	if err != nil {
		return RecipientId, RecipientUserId, err
	}

	for row.Next() {
		var a int
		var b int
		err = row.Scan(&a, &b)
		if err != nil {
			return RecipientId, RecipientUserId, err
		}
		RecipientId = a
		RecipientUserId = b
	}

	return RecipientId, RecipientUserId, nil
}
