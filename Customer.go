package main

type Customer struct {
	name string
}

func newCust(newName string) Customer {
	return Customer{
		name: newName,
	}
}
func toString(cust Customer) string {
	return cust.name
}
