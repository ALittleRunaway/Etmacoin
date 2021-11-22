package core

import (
	transaction "Blockchain/database/transaction"
	"strings"
)

func ProofOfWork(lastTransaction transaction.Transaction) (string, int, error) {
	var hash = ""
	for {
		lastTransaction.PoW += 1
		hash = CreateHash(lastTransaction)
		if strings.HasPrefix(hash, "000000") {
			return hash, lastTransaction.PoW, nil
		}
	}
}
