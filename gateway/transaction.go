package gateway

import (
	"Blockchain/database/transaction"
	"Blockchain/usecase"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var tpl_home = template.Must(template.ParseFiles("static/homepage/index.html"))

func HandlerHomePage(w http.ResponseWriter, r *http.Request) {
	tpl_home.Execute(w, nil)
}

func NewTransactionGateway(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(string(r.URL.Query()["user_id"][0]))
	recipientWallet := string(r.URL.Query()["recipient_wallet"][0])
	amount, _ := strconv.Atoi(string(r.URL.Query()["amount"][0]))
	message := string(r.URL.Query()["message"][0])

	newTransactionPlain := transaction.TransactionPlain{userId, recipientWallet, amount, message}
	newTransaction, err := usecase.AddNewTransactionUseCase(newTransactionPlain)
	if err != nil {
		fmt.Printf("Normal error message: %s", err.Error())
	}
	js, err := json.Marshal(newTransaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

