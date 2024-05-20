package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

var ErrInsuffientBalance = errors.New("insufficient balance")

func (w *Wallet) Deposit(amt Bitcoin) {
	w.balance += amt
}

func (w *Wallet) Withdraw(amt Bitcoin) error {
	if amt > w.balance {
		return ErrInsuffientBalance
	}
	w.balance -= amt
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
