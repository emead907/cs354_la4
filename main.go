package main

import . "./customer"
import . "./account"
import . "./bank"

import . "fmt"

//import . "./account"

func main() {
	var em = map[*IAccount]IAccount{}
	var bank = NewBank(em)
	var c = NewCust("Ann")
	var ca = NewChecking("01001", *c, 100.00)
	bank.Add(ca)
	var sa = NewSaving("01002", *c, 200.00)
	bank.Add(sa)

	//go bank.Accrue(0.02, accrue)
	Printf("%.2f\n", bank.Accrue(.02))
	Printf("%s", bank.String())

}
