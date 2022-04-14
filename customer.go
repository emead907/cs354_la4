package customer
type Customer struct {
	name string
}

func NewCust(newName string) (cus *Customer) {
	cus = new(Customer)
	cus.Init(newName)
	return
}
func (cus *Customer) Init(name string) {
	cus.name = name
}
func (cust *Customer) String() string {
	return cust.name
}
