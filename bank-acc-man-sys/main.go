package main

import (
	"fmt"
	"sync"
)

type Bank struct {
	Accounts []Account
}

type Account struct {
	Id      int
	Owner   string
	Balance float64
}

var accounts []Account

func printAccounts() {
	for _, acc := range accounts {
		fmt.Println(acc)
	}
}

func createAccount(id int, owner string, balance float64) *Account {
	accounts = append(accounts, Account{Id: id, Owner: owner, Balance: balance})
	return &accounts[len(accounts)-1]
}

var mu sync.Mutex

func deposit(acc *Account, amount float64) error {
	//lock to protect shared data
	mu.Lock()
	defer mu.Unlock()
	if amount <= 0 {
		return fmt.Errorf("deposit amount must be more than zero")
	}
	acc.Balance += amount
	return nil
}

func withdraw(ac *Account, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("withdraw amount must be more than zero")
	}
	if amount > ac.Balance {
		return fmt.Errorf("insufficient funds")
	}
	ac.Balance -= amount
	return nil
}

func findAccount(id int, accounts []Account) (*Account, error) {
	if len(accounts) == 0 {
		return nil, fmt.Errorf("no accounts available")
	}
	for i := 0; i < len(accounts); i++ {
		if accounts[i].Id == id {
			return &accounts[i], nil
		}
	}
	return nil, fmt.Errorf("account not found")
}

func transfer(from, to *Account, amount float64) error {
	if err := withdraw(from, amount); err != nil {
		return err
	}
	return deposit(to, amount)
}

func main() {
	fmt.Println("Bank Account Management System")
	acc := createAccount(1, "faze", 1000)
	fmt.Println("Created Account:", acc)
	acc2 := createAccount(2, "doe", 2000)
	fmt.Println("Created Account:", acc2)

	//wait for goroutines to finish
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			err := deposit(acc, 1)
			if err != nil {
				fmt.Println("Error during deposit:", err)
			} else {
				fmt.Println("After Deposit:", acc)
			}
		}()
	}
	wg.Wait()

	err := withdraw(acc2, 2500)
	if err != nil {
		fmt.Println("Error during withdrawal:", err)
	} else {
		fmt.Println("After Withdrawal:", acc2)
	}
	foundAcc, err := findAccount(1, accounts)
	if err != nil {
		fmt.Println("Error finding account:", err)
	} else {
		fmt.Println("Found Account:", *foundAcc)
	}

	err = transfer(acc, acc2, 1000)
	if err != nil {
		fmt.Println("Error during transfer:", err)
	} else {
		fmt.Println("After Transfer:")
		fmt.Println("From Account:", acc)
		fmt.Println("To Account:", acc2)
	}

	fmt.Println("All Accounts:")
	printAccounts()
}
