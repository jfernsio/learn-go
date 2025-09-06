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

var accounts []*Account //--> making slice of pointers to Account struct

func printAccounts() {
	mu.Lock()
	defer mu.Unlock()
	if len(accounts) == 0 {
		fmt.Println("No accounts available")
		return
	}
	for _, acc := range accounts {
		fmt.Println(acc)
	}
}

func createAccount(id int, owner string, balance float64) *Account {
	mu.Lock()
	defer mu.Unlock()
	if balance < 0 {
		balance = 0
	}
	accounts = append(accounts, &Account{Id: id, Owner: owner, Balance: balance})
	return accounts[len(accounts)-1]
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
	mu.Lock()
	defer mu.Unlock()
	if amount <= 0 {
		return fmt.Errorf("withdraw amount must be more than zero")
	}
	if amount > ac.Balance {
		return fmt.Errorf("insufficient funds")
	}
	ac.Balance -= amount
	return nil
}

func findAccount(id int) (*Account, error) {
	mu.Lock()
	defer mu.Unlock()
	if len(accounts) == 0 {
		return nil, fmt.Errorf("no accounts available")
	}
	for _, acc := range accounts {
		if acc.Id == id {
			return acc, nil
		}
	}
	return nil, fmt.Errorf("account not found")
}

func transfer(from, to *Account, amount float64) error {
	mu.Lock()
	defer mu.Unlock()

	if from.Balance < amount {
		return fmt.Errorf("insufficient funds")
	}
	from.Balance -= amount
	to.Balance += amount
	return nil

}

func main() {
	fmt.Println("Bank Account Management System")
	acc := createAccount(1, "faze", 1000)
	fmt.Println("Created Account:", acc)
	acc2 := createAccount(2, "doe", 2000)
	fmt.Println("Created Account:", acc2)

	//wait for goroutines to finish
	var wg sync.WaitGroup
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func() {
			defer wg.Done()
			err := deposit(acc, 1)
			if err != nil {
				fmt.Println("Error during deposit:", err)
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
	foundAcc, err := findAccount(1)
	if err != nil {
		fmt.Println("Error finding account:", err)
	} else {
		fmt.Println("Found Account:", *foundAcc)
	}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err = transfer(acc, acc2, 100)
			if err != nil {
				fmt.Println("Error during transfer of acc1:", err)
			} else {
				// fmt.Println("After Transfer Acc1 -> Acc2:")
				// fmt.Println("From Account:", acc)
				// fmt.Println("To Account:", acc2)
			}
			err2 := transfer(acc2, acc, 50)
			if err2 != nil {
				fmt.Println("Error during transfer of acc2:", err2)
			} else {
				// fmt.Println("After Transfer Acc2 -> Acc1:")
				// fmt.Println("From Account:", acc2)
				// fmt.Println("To Account:", acc)
			}
		}()
	}
	wg.Wait()

	fmt.Println("All Accounts:")
	printAccounts()
	// Calculate total money in system
	mu.Lock()
	defer mu.Unlock()
	total := 0.0
	for _, acc := range accounts {
		total += acc.Balance
	}
	fmt.Printf("Total money in system: $%.2f\n", total) //13000
	fmt.Printf("Expected: $%.2f\n", 13000.0)            // 3000 original + 100 deposits of $1

	// 10500
}
