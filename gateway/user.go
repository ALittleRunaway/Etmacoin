package gateway

import (
	"Blockchain/database/user"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
)

func HandleNewUser(w http.ResponseWriter, r *http.Request) {
	login := r.URL.Query()["login"]
	pass := r.URL.Query()["pass"]

	passEncrypted, err := Encrypt(string(pass[0]), MySecret)
	if err != nil {
		fmt.Println(err)
	}
	newUserPlain := database.UserPlain{Login: string(login[0]), Password: passEncrypted}
	newUser, err := database.AddNewUserHandler(newUserPlain)
	if err != nil {
		fmt.Println(err)
	}
	js, err := json.Marshal(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func HandleGetUserInfo(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(r.URL.Query()["user_id"][0])

	userInfo, err := database.GetUserInfoHandler(userId)
	if err != nil {
		fmt.Println(err)
	}
	js, err := json.Marshal(userInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func HandleLoginUser(w http.ResponseWriter, r *http.Request) {
	login := r.URL.Query()["login"]
	pass := r.URL.Query()["pass"]

	passEncrypted, err := Encrypt(string(pass[0]), MySecret)
	if err != nil {
		fmt.Println(err)
	}
	newUserPlain := database.UserPlain{Login: string(login[0]), Password: passEncrypted}

	userInfo, err := database.LoginUserHandler(newUserPlain)
	if err != nil {
		fmt.Println(err)
	}
	js, err := json.Marshal(userInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
