package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

type Wallet struct{
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
	fmt.Printf("address of balance in Deposit is %v\n", &w.balance)
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdrawal(amount Bitcoin) error {
	if amount > w.balance {
		//can create a message of your choice
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}