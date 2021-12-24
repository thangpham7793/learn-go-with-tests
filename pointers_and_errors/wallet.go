package pointers_and_errors

import (
	"errors"
	"fmt"
)

type Bitcoin int

// global to the package
var InsufficientFundError = errors.New("insufficient fund")

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	// auto dereference (==  (*w).balance)
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return InsufficientFundError
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// this can then be used as %s in print func
// implicitly implements fmt.Stringer()
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
