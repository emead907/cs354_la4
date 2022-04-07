package main

type Customer struct {
	name string
}

func newCust(newName string) Customer {
	return Customer{
		name: newName,
	}
}
func toStringC(cust Customer) string {
	return cust.name
}
