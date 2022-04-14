package main

import . "fmt"

//import . "./account"

func main() {
	var em = map[*IAccount]IAccount{}
	var bank = newBank(em)
	var c = newCust("Ann")
	var ca = newChecking("01001", *c, 100.00)
	bank.Add(ca)
	var sa = newSaving("01002", *c, 200.00)
	bank.Add(sa)
	var c1 = newCust("Emily")
	var sa1 = newSaving("01003", *c1, 200.00)
	bank.Add(sa1)

	//go bank.Accrue(0.02, accrue)
	Printf("%.2f\n", bank.Accrue(.02))
	Printf("%s", bank.String())

}

/*
public static void main(String[] args) {
	bank.add(new CheckingAccount("01001",c,100.00));
	bank.add(new SavingAccount("01002",c,200.00));
	bank.accrue(0.02);
	System.out.println(bank);
    }
*/
