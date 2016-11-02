// Package account bank...
package account

import "sync"

// Account struct
type Account struct {
	Open    bool
	balance int64
	sync.Mutex
}

// Open an account with a deposit
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{Open: true, balance: initialDeposit}
}

// Close an account and payout remaining balance
func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	a.Open = false
	payout = a.balance
	a.balance = 0
	a.Unlock()
	ok = true
	return
}

// Balance on account
func (a *Account) Balance() (balance int64, ok bool) {
	a.Lock()
	if !a.Open {
		return 0, false
	}
	balance = a.balance
	a.Unlock()
	ok = true
	return
}

// Deposit (or withdrawal if negative amount)
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.Lock()
	if a.balance+amount < 0 || !a.Open {
		return
	}
	a.balance += amount
	newBalance = a.balance
	a.Unlock()
	ok = true
	return
}
