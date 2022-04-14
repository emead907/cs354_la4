package account

type IChecking interface {
	IAccount
}

type Checking struct {
	Account
}

func (c *Checking) Accure(rate float32, accrue chan float32) {
	c.Account.Accrue(rate, accrue)
}

func newChecking(number string, customer Customer, balance float32) (c *Checking) {
	c = new(Checking)
	c.Init(number, customer, balance)
	return
}

func (c *Checking) Init(number string, customer Customer, balance float32) {
	c.Account.Init(number, customer, balance)
}
