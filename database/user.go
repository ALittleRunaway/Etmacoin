package database

import (
	"github.com/google/uuid"
)

type UserPlain struct {
	Login    string
	Password string
}

type User struct {
	UserPlain
	Id      int
	Balance int
}

func AddNewUser(newUser UserPlain) error {
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

	if _, err = queryConn.Exec(newUser.Login, newUser.Password, uuid.New(), 100); err != nil {
		return err
	}

	return nil
}
