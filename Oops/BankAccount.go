package main

import (
	"errors"
	"fmt"
)

type Account struct {
	AccountNumber string
	Balance       float64
	AccountHolder string
}

type SavingAccount struct {
	Account
	InterestRate	float64
}

type LoanAccount struct {
	Account
	LoanAmount	float64
}

func (acc *Account) Deposite(amount float64) error {
	if amount <= 0 {
		return errors.New("Amount much be positive")
	}
	acc.Balance += amount
	fmt.Printf("Deposited: $%.2f to Account no. ending with XXXX-XXXX-%s. New Balance: $%.2f\n",
	 amount, acc.AccountNumber[4:], acc.Balance)
	return nil
}

func (acc* Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("Amount much be positive")
	}
	if(amount > acc.Balance) {
		return errors.New("Insufficient Balance! Cannot withdraw more than what you have")
	}
	acc.Balance -= amount
	fmt.Printf("Withdrawal of $%.2f form Account no. ending with XXXX-XXXX-%s is successfull. New Balance: $%.2f\n",
	 amount, acc.AccountNumber[4:], acc.Balance)
	return nil
}

func (acc Account) GetBalance() float64 { return acc.Balance}

func (acc Account) String() string {
	return fmt.Sprintf("---Account Details---\nAccount Holder: %s\nAccount Number: %s\nAccount Balance: $%.2f",
	 acc.AccountHolder, acc.AccountNumber, acc.Balance)
}

func (acc* SavingAccount) AddInterest() error {
	interest := acc.GetBalance() * acc.InterestRate
	err := acc.Deposite(interest)
	if(err != nil) {
		return err
	}
	fmt.Printf("Adding interest $%.2f to saving account ending with %s\n. New Balance: %.2f\n", 
		interest, acc.AccountNumber[4:], acc.GetBalance())
	return nil
}

func (acc* LoanAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("Amount much be positive")
	}
	if(amount > acc.Balance + acc.LoanAmount) {
		return errors.New("Insufficient Balance! Cannot withdraw more than what you have")
	}

	if(acc.GetBalance() <= amount) {
		acc.LoanAmount -= amount - acc.Balance
		acc.Balance = 0
	}else{
		acc.Balance -= amount
	}

	fmt.Printf("Withdrawal of $%.2f form Account no. ending with XXXX-XXXX-%s is successfull\n. New Acc Balance: $%.2f New Loan Balance: $%.2f\n",
		amount, acc.AccountNumber[4:], acc.Balance, acc.LoanAmount)
	return nil
}

func main() {
	// Note: AccountNumber needs at least 5 characters because [4:] slices from index 4.
	savings := SavingAccount{
		Account: Account{
			AccountNumber: "12345678",
			Balance:       1000.00,
			AccountHolder: "Navneet Ranjan",
		},
		InterestRate: 0.05, // 5%
	}

	loan := LoanAccount{
		Account: Account{
			AccountNumber: "87654321",
			Balance:       500.00,
			AccountHolder: "Navneet Ranjan",
		},
		LoanAmount: 2000.00,
	}

	fmt.Println("=== Initial Account Details (via String()) ===")
	fmt.Println(savings)
	fmt.Println(loan)

	fmt.Println("\n=== SavingsAccount Operations ===")
	if err := savings.Deposite(250); err != nil {
		fmt.Println("Error:", err)
	}
	if err := savings.Withdraw(300); err != nil {
		fmt.Println("Error:", err)
	}
	if err := savings.AddInterest(); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Savings GetBalance(): $%.2f\n", savings.GetBalance())

	fmt.Println("\n=== LoanAccount Operations ===")
	if err := loan.Deposite(200); err != nil {
		fmt.Println("Error:", err)
	}
	// Withdraws more than Balance (700), dipping into LoanAmount
	if err := loan.Withdraw(1000); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Loan GetBalance(): $%.2f\n", loan.GetBalance())

	fmt.Println("\n=== Error Handling Checks ===")
	if err := savings.Deposite(-50); err != nil {
		fmt.Println("Expected deposit error:", err)
	}
	if err := savings.Withdraw(50000); err != nil {
		fmt.Println("Expected withdraw error:", err)
	}
	if err := loan.Withdraw(50000); err != nil {
		fmt.Println("Expected loan withdraw error:", err)
	}
}