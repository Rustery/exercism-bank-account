package account

import "sync"

// Define the Account type here.
type Account struct {
	amount int64
	closed bool
	m      sync.RWMutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{amount: amount}
}

func (a *Account) Balance() (int64, bool) {
	a.m.RLock()
	defer a.m.RUnlock()
	return a.amount, !a.closed
}

func (a *Account) Deposit(amount int64) (result int64, success bool) {
	a.m.Lock()
	defer a.m.Unlock()
	if a.closed || amount < 0 && a.amount+amount < 0 {
		return a.amount, false
	}
	a.amount += amount
	return a.amount, true
}

func (a *Account) Close() (int64, bool) {
	a.m.Lock()
	defer a.m.Unlock()
	if a.closed {
		return 0, false
	}
	amount := a.amount
	a.amount, a.closed = 0, true
	return amount, true
}
