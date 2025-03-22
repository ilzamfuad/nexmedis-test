package main

import (
	"errors"
	"sync"
)

type BankAccount struct {
	balance int
	mutex   sync.Mutex
}

func (a *BankAccount) Deposit(amount int) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	a.balance += amount
}

func (a *BankAccount) Withdraw(amount int) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if amount > a.balance {
		return errors.New("insufficient balance")
	}

	a.balance -= amount
	return nil
}

func (a *BankAccount) Balance() int {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	return a.balance
}
