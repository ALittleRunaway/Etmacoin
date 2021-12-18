package cryptocore

import (
	"Blockchain/database/transaction"
	"strings"
)

func ProofOfWork(transaction transaction.Transaction) (int, error) {
	var hash = ""
	for {
		transaction.PoW += 1
		hash = CreateHash(transaction)
		if strings.HasPrefix(hash, "6666") {
			return transaction.PoW, nil
		}
	}
}
