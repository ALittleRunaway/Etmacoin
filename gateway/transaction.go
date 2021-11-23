package gateway

import (
	"Blockchain/core"
	"Blockchain/database"
	transaction "Blockchain/database/transaction"
	user "Blockchain/database/user"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func HandleNewTransaction(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(string(r.URL.Query()["user_id"][0]))
	recipientWallet := string(r.URL.Query()["recipient_wallet"][0])
	amount, _ := strconv.Atoi(string(r.URL.Query()["amount"][0]))
	message := string(r.URL.Query()["message"][0])

	newTransactionPlain := transaction.TransactionPlain{userId, recipientWallet, amount, message}
	newTransaction, err := AddNewTransactionHandler(newTransactionPlain)
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

func AddNewTransactionHandler(newTransactionPlain transaction.TransactionPlain) (transaction.AddNewTransactionResponse, error) {
	db, err := database.Connection()
	var newTransaction transaction.Transaction
	var newTransactionResponse transaction.AddNewTransactionResponse
	newTransaction.Amount = newTransactionPlain.Amount
	newTransaction.Message = newTransactionPlain.Message
	newTransaction.SenderUserId = newTransactionPlain.UserId
	newTransaction.Time = time.Now().Add(time.Hour * 3)

	if err != nil {
		newTransactionResponse.Response = err.Error()
		return newTransactionResponse, err
	}
	newTransaction.SenderId, err = user.GetSenderId(db, newTransactionPlain.UserId)
	if err != nil {
		newTransactionResponse.Response = err.Error()
		return newTransactionResponse, err
	}
	newTransaction.RecipientId, newTransaction.RecipientUserId, err =
		user.GetRecipientAndUserId(db, newTransactionPlain.RecipientWallet)
	if err != nil {
		newTransactionResponse.Response = err.Error()
		return newTransactionResponse, err
	}
	lastTransaction, err := transaction.GetLastTransaction(db)
	if err != nil {
		newTransactionResponse.Response = err.Error()
		return newTransactionResponse, err
	}
	newTransaction.PrevHash, newTransaction.PoW, err = core.ProofOfWork(lastTransaction)
	if err != nil {
		newTransactionResponse.Response = err.Error()
		return newTransactionResponse, err
	}
	err = transaction.AddNewTransaction(db, newTransaction)
	if err != nil {
		newTransactionResponse.Response = err.Error()
		return newTransactionResponse, err
	}
	newTransactionResponse.Transaction = newTransaction
	err = transaction.TakeCoinsFromSender(db, newTransaction)
	if err != nil {
		newTransactionResponse.Response = err.Error()
		return newTransactionResponse, err
	}
	err = transaction.AddCoinsToRecipient(db, newTransaction)
	if err != nil {
		newTransactionResponse.Response = err.Error()
		return newTransactionResponse, err
	}
	newTransactionResponse.Response = "The transaction was sent and mined successfully!"
	return newTransactionResponse, nil
}
