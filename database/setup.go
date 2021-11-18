package database

import (
	"database/sql"
)

func setup() {
	db, err := sql.Open("mysql", "root:Everything7tays@tcp(127.0.0.1:3306)/blockchain")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}
