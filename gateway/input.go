package gateway

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
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

func HandleNewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query())
	login := r.URL.Query()["login"]
	pass := r.URL.Query()["pass"]

	log.Println("Url Param 'login' is: " + string(login[0]))
	log.Println("Url Param 'pass' is: " + string(pass[0]))
	//login := r.URL.Query()["login"]
	//pass := r.URL.Query()["pass"]
	//fmt.Println(string(login[0]))
	//fmt.Println(string(pass[0]))

	passEncrypted, err := Encrypt(string(pass[0]), MySecret)
	if err != nil {
		fmt.Println(err)
	}
	newUser := UserPlain{Login: string(login[0]), Password: passEncrypted}
	err = AddNewUser(newUser)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newUser)
}

//
//func HandleNewUser(writer http.ResponseWriter, r *http.Request) {
//	login := string(r.URL.Query()["login"][0])
//	pass := string(r.URL.Query()["pass"][0])
//	fmt.Println(login)
//
//	passEncrypted, err := Encrypt(pass, MySecret)
//	if err != nil {
//		fmt.Println(err)
//	}
//	newUser := UserPlain{Login: login, Password: passEncrypted}
//	err = AddNewUser(newUser)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(newUser)
//}
//
//func HandleNewUser(writer http.ResponseWriter, r *http.Request) {
//	login := string(r.URL.Query()["login"][0])
//	pass := string(r.URL.Query()["pass"][0])
//
//	passEncrypted, err := Encrypt(pass, MySecret)
//	if err != nil {
//		fmt.Println(err)
//	}
//	newUser := UserPlain{Login: login, Password: passEncrypted}
//	err = AddNewUser(newUser)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(newUser)
//}
//	//passDecrypted, err := Decrypt(passEncrypted, MySecret)
//	//if err != nil {
//	//	fmt.Println("error decrypting your encrypted text: ", err)
//	//}
//	//fmt.Println(login)
//	//fmt.Println(pass)
//	//fmt.Println(passEncrypted)
//	//fmt.Println(passDecrypted)
//	return nil

func AddNewUser(newUser UserPlain) error {
	db, err := sql.Open("mysql", "root:Everything7tays@tcp(127.0.0.1:3306)/blockchain")
	if err != nil {
		panic(err.Error())
		return err
	}
	//defer db.Close()
	insert, err := db.Query("INSERT INTO blockchain.user (login, password, balance) " +
		"VALUES ('" + newUser.Login + "', '" + newUser.Password + "', 100);")
	if err != nil {
		panic(err.Error())
		return err
	}
	defer insert.Close()
	return nil
}
