package account

import (
	. "fmt"
)

type IAccount interface {
	Accure(rate float32, accrue chan float32)
	Balance() float32
	Deposit(amount float32)
	Withdraw(amount float32)
	String() string
}

type Account struct {
	number   string
	customer Customer
	balance  float32
}

func NewAccount(number string, customer Customer, balance float32) (a *Account) {
	a = new(Account)
	a.Init(number, customer, balance)
	return
}

func (a *Account) Init(number string, customer Customer, balance float32) {
	a.balance = balance
	a.number = number
	a.customer = customer
}

func (a *Account) Accrue(rate float32, accrue chan float32) {
	accrue <- 0
}
func (a *Account) Balance() float32 {
	return a.balance
}

func (a *Account) Withdraw(amount float32) {
	a.balance = a.balance - amount
}
func (a *Account) Deposit(amount float32) {
	a.balance = a.balance + amount
}
func (a *Account) String() string {
	return a.number + ":" + a.customer.String() + ":" + Sprintf("%.2f", a.balance)
}
