package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type DbParams struct {
	User     string
	Password string
	Host     string
	Port     int
	Name     string
	Driver   string
}

func Connection() (conn *sql.DB, err error) {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbParams := DbParams{
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		Name:     os.Getenv("DB_NAME"),
		Driver:   os.Getenv("DB_DRIVER_NAME"),
	}

	db, err := sql.Open(dbParams.Driver, fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true", dbParams.User, dbParams.Password, dbParams.Host, dbParams.Port, dbParams.Name))
	if err != nil {
		return nil, errors.New("error during the connection to the database")
	}

	return db, nil
}
