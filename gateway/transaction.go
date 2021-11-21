package gateway

import (
	"Blockchain/database/transaction"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func HandleNewTransaction(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(string(r.URL.Query()["user_id"][0]))
	recipientWallet := string(r.URL.Query()["recipient_wallet"][0])
	amount, _ := strconv.Atoi(string(r.URL.Query()["amount"][0]))
	message := string(r.URL.Query()["message"][0])

	newTransactionPlain := database.TransactionPlain{userId, recipientWallet, amount, message}
	newTransaction, err := database.AddNewTransactionHandler(newTransactionPlain)
	if err != nil {
		fmt.Println(err)
	}
	js, err := json.Marshal(newTransaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
