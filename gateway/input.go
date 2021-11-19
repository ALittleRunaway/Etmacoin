package gateway

import (
	"Blockchain/database"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func HandleNewUser(w http.ResponseWriter, r *http.Request) {
	login := r.URL.Query()["login"]
	pass := r.URL.Query()["pass"]

	passEncrypted, err := Encrypt(string(pass[0]), MySecret)
	if err != nil {
		fmt.Println(err)
	}
	newUser := database.UserPlain{Login: string(login[0]), Password: passEncrypted}
	err = database.AddNewUser(newUser)
	if err != nil {
		fmt.Println(err)
	}
}
