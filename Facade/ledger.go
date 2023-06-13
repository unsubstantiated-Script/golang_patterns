package Facade

import "fmt"

type Ledger struct {
}

func (s *Ledger) makeEntry(accountID, txnType string, amount int) {
	fmt.Printf("Make ledger entry for accountID %s with txnType %s for amount %d\n", accountID, txnType, amount)
	return
}
