package gateway

import (
	"Blockchain/database/transaction"
	"Blockchain/usecase"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func NewTransactionGateway(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(string(r.URL.Query()["user_id"][0]))
	recipientWallet := string(r.URL.Query()["recipient_wallet"][0])
	amount, _ := strconv.Atoi(string(r.URL.Query()["amount"][0]))
	message := string(r.URL.Query()["message"][0])

	newTransactionPlain := transaction.TransactionPlain{userId, recipientWallet, amount, message}
	newTransaction, err := usecase.AddNewTransactionUseCase(newTransactionPlain)
	if err != nil {
		log.Printf("New transaction wasn't added. SenderId: %d, RecipientId: %d, Amount: %d, Message: %s, Time: %s",
			newTransaction.SenderUserId, newTransaction.RecipientUserId, newTransaction.Amount, newTransaction.Message, newTransaction.Time)
		log.Println(err)
	}
	log.Printf("New transaction was added sucessfully. SenderId: %d, RecipientId: %d, Amount: %d, Message: %s, Time: %s",
		newTransaction.SenderUserId, newTransaction.RecipientUserId, newTransaction.Amount, newTransaction.Message, newTransaction.Time)
	js, err := json.Marshal(newTransaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func GetLatestTransactionsGateway(w http.ResponseWriter, r *http.Request) {
	latestTransactions, err := usecase.GetLatestTransactionsUseCase()
	if err != nil {
		fmt.Printf("Normal error message: %s", err.Error())
	}
	js, err := json.Marshal(latestTransactions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func GetAllTransactionsGateway(w http.ResponseWriter, r *http.Request) {
	allTransactions, err := usecase.GetAllTransactionsUseCase()
	if err != nil {
		fmt.Printf("Normal error message: %s", err.Error())
	}
	js, err := json.Marshal(allTransactions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func GetUserTransactionsGateway(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(string(r.URL.Query()["user_id"][0]))
	userTransactions, err := usecase.GetUserTransactionsUseCase(userId)
	if err != nil {
		fmt.Printf("Normal error message: %s", err.Error())
	}
	js, err := json.Marshal(userTransactions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
