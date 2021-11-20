package database

import (
	"database/sql"
	"github.com/google/uuid"
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

func AddNewUserHandler(newUserPlain UserPlain) (User, error) {
	db, err := Connection()
	var newUser User
	if err != nil {
		return newUser, err
	}
	wallet := uuid.New().String()
	err = AddNewUser(db, newUserPlain, wallet)
	if err != nil {
		return newUser, err
	}
	userId, err := GetUserId(db, wallet)
	if err != nil {
		return newUser, err
	}
	err = AddNewSender(db, userId)
	if err != nil {
		return newUser, err
	}
	err = AddNewRecipient(db, userId)
	if err != nil {
		return newUser, err
	}
	newUser = User{newUserPlain, userId, wallet, 100}
	return newUser, nil
}

func AddNewUser(db *sql.DB, newUser UserPlain, wallet string) error {
	db, err := Connection()
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
