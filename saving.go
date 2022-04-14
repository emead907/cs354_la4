package account

import . "./customer"

type ISaving interface {
	IAccount
}

type Saving struct {
	Account
	interest float32
}

func (s *Saving) Accure(rate float32, accrue chan float32) {
	interest := s.balance * rate
	s.balance = s.balance + (s.balance * rate)
	accrue <- interest
}

func newSaving(number string, customer Customer, balance float32) (s *Saving) {
	s = new(Saving)
	s.Init(number, customer, balance)
	return
}

func (s *Saving) Init(number string, customer Customer, balance float32) {
	s.Account.Init(number, customer, balance)
	s.interest = 0
}
