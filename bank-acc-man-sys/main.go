package main

import ("fmt")



type Account struct {
	Id int
	Owner string
	Balance float64
}

var accounts []Account

func createAccount(id int, owner string, balance float64) Account {
	accounts = append(accounts, Account{Id: id, Owner: owner, Balance: balance})
	return accounts[len(accounts)-1]
}

func deposit(acc *Account, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("deposit amount must be more than zero")
	}
	acc.Balance += amount
	return nil
}

func withdraw(ac * Account, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("withdraw amount must be more than zero")
	}
	if amount > ac.Balance {
		return fmt.Errorf("insufficient funds")
	}
	ac.Balance -= amount
	return nil
}







func main () {
	fmt.Println("Bank Account Management System")
	acc :=	createAccount(1,"faze", 1000)
	fmt.Println("Created Account:", acc)
	acc2 := createAccount(2,"doe", 2000)
	fmt.Println("Created Account:", acc2)
	err := deposit(&acc,500)
	if err != nil {
		fmt.Println("Error during deposit:", err)
	} else {
		fmt.Println("After Deposit:", acc)
	}
	err = withdraw(&acc2, 2500)
	if err != nil {
		fmt.Println("Error during withdrawal:", err)
	} else {
		fmt.Println("After Withdrawal:", acc2)
	}
}
