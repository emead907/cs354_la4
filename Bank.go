package bank

import ."./account"

type Bank struct {
	accounts map[*IAccount]IAccount
}

func newBank(accounts map[*IAccount]IAccount) (bank *Bank) {
	bank = new(Bank)
	bank.Init(accounts)
	return
}

func (bank *Bank) Init(accounts map[*IAccount]IAccount) {
	bank.accounts = accounts
}

func (bank *Bank) Add(account IAccount) {
	bank.accounts[&account] = account
}

func (bank *Bank) Accrue(rate float32) float32 {
	accrue := make(chan float32)
	var number float32
	number = 0
	for _, a := range bank.accounts {
		go a.Accure(rate, accrue)
	}
	for range bank.accounts {
		number += <-accrue
	}
	return number
}

func (bank *Bank) String() string {
	var ret = ""
	for b, a := range bank.accounts {
		b = b
		ret = ret + a.String() + "\n"
	}
	return ret
}
