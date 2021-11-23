package usecase

import (
	"Blockchain/cryptocore"
	"Blockchain/database"
	transaction "Blockchain/database/transaction"
	user "Blockchain/database/user"
	"time"
)

func AddNewTransactionUseCase(newTransactionPlain transaction.TransactionPlain) (transaction.AddNewTransactionResponse, error) {
	var newTransaction transaction.Transaction
	var newTransactionResponse transaction.AddNewTransactionResponse
	db, err := database.Connection()
	if err != nil {
		return newTransactionResponse, err
	}
	newTransaction.Amount = newTransactionPlain.Amount
	newTransaction.Message = newTransactionPlain.Message
	newTransaction.SenderUserId = newTransactionPlain.UserId
	newTransaction.Time = time.Now().Add(time.Hour * 3)

	newTransaction.SenderId, err = user.GetSenderId(db, newTransactionPlain.UserId)
	if err != nil {
		newTransactionResponse.Response = "Internal server error! Please, contact the developer."
		return newTransactionResponse, err
	}
	newTransaction.RecipientId, newTransaction.RecipientUserId, err =
		user.GetRecipientAndUserId(db, newTransactionPlain.RecipientWallet)
	if err != nil || newTransaction.RecipientId == 0 || newTransaction.RecipientUserId == 0{
		newTransactionResponse.Response = "There is no user with wallet id like this: " +
			"" + newTransactionPlain.RecipientWallet
		return newTransactionResponse, err
	}
	lastTransaction, err := transaction.GetLastTransaction(db)
	if err != nil {
		newTransactionResponse.Response = "Internal server error! Please, contact the developer."
		return newTransactionResponse, err
	}
	newTransaction.PrevHash, newTransaction.PoW, err = cryptocore.ProofOfWork(lastTransaction)
	if err != nil {
		newTransactionResponse.Response = "Internal server error! Please, contact the developer."
		return newTransactionResponse, err
	}
	err = transaction.AddNewTransaction(db, newTransaction)
	if err != nil {
		newTransactionResponse.Response = "Internal server error! Please, contact the developer."
		return newTransactionResponse, err
	}
	newTransactionResponse.Transaction = newTransaction
	err = transaction.TakeCoinsFromSender(db, newTransaction)
	if err != nil {
		newTransactionResponse.Response = "Internal server error! Please, contact the developer."
		return newTransactionResponse, err
	}
	err = transaction.AddCoinsToRecipient(db, newTransaction)
	if err != nil {
		newTransactionResponse.Response = "Internal server error! Please, contact the developer."
		return newTransactionResponse, err
	}
	userInfo, err := user.GetUserInfo(db, newTransaction.RecipientUserId)
	if err != nil {
		return newTransactionResponse, err
	}
	if newTransactionResponse.Response == "" {
		newTransactionResponse.Response = "The transaction to " + userInfo.Login + " was sent and mined successfully!"
	}
	return newTransactionResponse, nil
}
