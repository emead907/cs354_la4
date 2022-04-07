package main

type Account interface {
	balance() float32
	deposit()
	withdraw()
	toString() string
}
type AbAccount struct {
	Account
	num  string
	bal  float32
	cust Customer
}

func toStringA(accountName string) string {
	return accountName
}

type CheckingAccount struct{ AbAccount }

func (c CheckingAccount) balance() float32 {
	return c.bal
}

func newChecking() *CheckingAccount {
	checking := CheckingAccount{AbAccount{}}
	checking.AbAccount.Account = checking
	return &checking
}

type SavingAccount struct{ Account }
