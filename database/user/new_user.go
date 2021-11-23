package user

import (
	"Blockchain/database"
	"database/sql"
)

type UserPlain struct {
	Login    string
	Password string
}

type User struct {
	UserPlain
	Id      int
	Wallet  string
	Balance int
}

func AddNewUser(db *sql.DB, newUser UserPlain, wallet string) error {
	db, err := database.Connection()
	if err != nil {
		return err
	}

	const query = `INSERT INTO blockchain.user (login, password, wallet, balance) VALUES (?, ?, ?, ?)`

	queryConn, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer queryConn.Close()

	if _, err = queryConn.Exec(newUser.Login, newUser.Password, wallet, 100); err != nil {
		return err
	}
	return nil
}

func GetUserId(db *sql.DB, wallet string) (int, error) {
	const query = `SELECT id FROM blockchain.user WHERE wallet = ?`
	userId := 0

	row, err := db.Query(query, wallet)
	if err != nil {
		return userId, err
	}

	for row.Next() {
		var i int
		err = row.Scan(&i)
		if err != nil {
			return userId, err
		}
		userId = i
	}

	return userId, nil
}

func AddNewSender(db *sql.DB, userId int) error {
	const query = `INSERT INTO blockchain.sender (user_id) VALUES (?)`

	queryConn, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer queryConn.Close()

	if _, err = queryConn.Exec(userId); err != nil {
		return err
	}
	return nil
}

func AddNewRecipient(db *sql.DB, userId int) error {
	const query = `INSERT INTO blockchain.recipient (user_id) VALUES (?)`

	queryConn, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer queryConn.Close()

	if _, err = queryConn.Exec(userId); err != nil {
		return err
	}
	return nil
}
