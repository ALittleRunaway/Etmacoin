package gateway

import (
	"Blockchain/cryptocore"
	"Blockchain/database/user"
	"Blockchain/usecase"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var tpl_login = template.Must(template.ParseFiles("static/login_page/index.html"))

func HandlerLoginPage(w http.ResponseWriter, r *http.Request) {
	tpl_login.Execute(w, nil)
}

func GetUserInfoGateway(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(r.URL.Query()["user_id"][0])

	userInfo, err := usecase.GetUserInfoUseCase(userId)
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


func NewUserGateway(w http.ResponseWriter, r *http.Request) {
	login := r.URL.Query()["login"]
	pass := r.URL.Query()["pass"]

	passEncrypted, err := cryptocore.Encrypt(string(pass[0]), cryptocore.MySecret)
	if err != nil {
		fmt.Println(err)
	}
	newUserPlain := user.UserPlain{Login: string(login[0]), Password: passEncrypted}
	newUser, err := usecase.AddNewUserUseCase(newUserPlain)
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


func LoginUserGateway(w http.ResponseWriter, r *http.Request) {
	login := r.URL.Query()["login"]
	pass := r.URL.Query()["pass"]

	passEncrypted, err := cryptocore.Encrypt(string(pass[0]), cryptocore.MySecret)
	if err != nil {
		fmt.Println(err)
	}
	newUserPlain := user.UserPlain{Login: string(login[0]), Password: passEncrypted}

	userInfo, err := usecase.LoginUserUseCase(newUserPlain)
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

func RandomWalletGateway(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(string(r.URL.Query()["user_id"][0]))

	randomWallet, err := usecase.RandomWalletUseCase(userId)
	if err != nil {
		fmt.Println(err)
	}
	js, err := json.Marshal(randomWallet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
