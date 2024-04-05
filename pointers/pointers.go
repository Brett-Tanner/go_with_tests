package main

import (
	"errors"
	"fmt"
)

var errInsufficientFunds = errors.New("Insufficient Funds")

type Stringer interface {
	String() string
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errInsufficientFunds
	}

	w.balance -= amount
	return nil
}
