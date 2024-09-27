package main

import (
	"errors"
	"fmt"
)

type Bitcoin float64

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%g BTC", b)
}

type Wallet struct {
	Amount Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.Amount += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.Amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.Amount < amount {
		return errors.New("Not enoth BTC")
	}
	w.Amount -= amount
	return nil
}

func main() {
	wallet := Wallet{Bitcoin(10.0)}

	err := wallet.Withdraw(Bitcoin(50))

	if err != nil {
		fmt.Println(err)
	}
}
