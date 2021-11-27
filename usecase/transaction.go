package usecase

import (
	"Blockchain/cryptocore"
	"Blockchain/database"
	"Blockchain/database/transaction"
	"Blockchain/database/user"
	"time"
)

func AddNewTransactionUseCase(newTransactionPlain transaction.TransactionPlain) (transaction.AddNewTransactionResponse, error) {
	var newTransaction transaction.TransactionExtend
	var newTransactionResponse transaction.AddNewTransactionResponse
	db, err := database.Connection()
	if err != nil {
		return newTransactionResponse, err
	}
	newTransaction.Amount = newTransactionPlain.Amount
	newTransaction.Message = newTransactionPlain.Message
	newTransaction.SenderUserId = newTransactionPlain.UserId
	newTransaction.Time = time.Now().Add(time.Hour * 3).Round(time.Duration(time.Second)).UTC()

	newTransaction.SenderId, err = user.GetSenderId(db, newTransactionPlain.UserId)
	if err != nil {
		newTransactionResponse.Response = "Internal server error! Please, contact the developer."
		return newTransactionResponse, err
	}
	newTransaction.RecipientId, newTransaction.RecipientUserId, err =
		user.GetRecipientAndUserId(db, newTransactionPlain.RecipientWallet)
	if err != nil || newTransaction.RecipientId == 0 || newTransaction.RecipientUserId == 0 {
		newTransactionResponse.Response = "There is no user with wallet id like this: " +
			"" + newTransactionPlain.RecipientWallet
		return newTransactionResponse, err
	}
	lastTransaction, err := transaction.GetLastTransaction(db)
	if err != nil {
		newTransactionResponse.Response = "Internal server error! Please, contact the developer."
		return newTransactionResponse, err
	}
	newTransaction.PrevHash = cryptocore.CreateHash(lastTransaction)
	var transactionBasic = transaction.Transaction{
		Id:          lastTransaction.Id + 1,
		SenderId:    newTransaction.SenderId,
		RecipientId: newTransaction.RecipientId,
		Amount:      newTransaction.Amount,
		Message:     newTransaction.Message,
		Time:        newTransaction.Time,
		PrevHash:    newTransaction.PrevHash,
		PoW:         0,
	}
	newTransaction.PoW, err = cryptocore.ProofOfWork(transactionBasic)
	transactionBasic.PoW = newTransaction.PoW
	if err != nil {
		newTransactionResponse.Response = "Internal server error! Please, contact the developer."
		return newTransactionResponse, err
	}
	err = transaction.AddNewTransaction(db, newTransaction)
	if err != nil {
		newTransactionResponse.Response = "Internal server error! Please, contact the developer."
		return newTransactionResponse, err
	}
	newTransactionResponse.TransactionExtend = newTransaction
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

func GetLatestTransactionsUseCase() (transaction.LatestTransactionsResponse, error) {
	var latestTransactionsResponse transaction.LatestTransactionsResponse
	db, err := database.Connection()
	if err != nil {
		return latestTransactionsResponse, err
	}
	latestTransactions, err := transaction.GetLatestTransactions(db)
	if err != nil {
		return latestTransactionsResponse, err
	}
	for _, lt := range latestTransactions {
		latestTransactionsResponse.Transactions = append(latestTransactionsResponse.Transactions,
			transaction.LatestTransactions{
				Hash:   cryptocore.CreateHash(lt),
				Time:   lt.Time,
				Amount: lt.Amount,
			})
	}
	return latestTransactionsResponse, nil
}

func GetAllTransactionsUseCase() (transaction.AllTransactionsResponse, error) {
	var allTransactionsResponse transaction.AllTransactionsResponse
	db, err := database.Connection()
	if err != nil {
		return allTransactionsResponse, err
	}
	allTransactions, err := transaction.GetAllTransactions(db)
	if err != nil {
		return allTransactionsResponse, err
	}
	for _, lt := range allTransactions {
		allTransactionsResponse.Transactions = append(allTransactionsResponse.Transactions,
			transaction.Transaction{
				Id:          lt.Id,
				SenderId:    lt.SenderId,
				RecipientId: lt.RecipientId,
				Amount:      lt.Amount,
				Message:     lt.Message,
				Time:        lt.Time,
				PrevHash:    lt.PrevHash,
				PoW:         lt.PoW,
			})
	}
	allTransactionsResponse.Count = len(allTransactions)
	return allTransactionsResponse, nil
}

func GetUserTransactionsUseCase(userId int) (transaction.UserTransactionsResponse, error) {
	var userTransactionsResponse transaction.UserTransactionsResponse
	db, err := database.Connection()
	if err != nil {
		return userTransactionsResponse, err
	}
	userTransactions, err := transaction.GetUserTransactions(db, userId)
	if err != nil {
		return userTransactionsResponse, err
	}
	for _, ut := range userTransactions {
		userTransactionsResponse.Transactions = append(userTransactionsResponse.Transactions,
			transaction.UserTransaction{
				CallerWallet: ut.CallerWallet,
				Amount:       ut.Amount,
				Message:      ut.Message,
				Time:         ut.Time,
				Direction:    ut.Direction,
			})
	}
	userTransactionsResponse.Count = len(userTransactions)
	return userTransactionsResponse, nil
}
