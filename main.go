package main

import (
	"fmt"
	"os"
)

func main() {
	var myExpenses Expenses
	storage := newStorage[Expenses]("Expenses.json")

	// Load expenses from the file
	err := storage.Load(&myExpenses)
	if err != nil {
		if os.IsNotExist(err) {
			// Initialize empty expenses slice
			myExpenses = make(Expenses, 0)
		} else {
			fmt.Println("Error loading expenses:", err)
			os.Exit(1)
		}
	}

	cmdFlags := newCommandFlags()
	if err := cmdFlags.Execute(&myExpenses); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Save the updated expenses to the file
	if err := storage.Save(myExpenses); err != nil {
		fmt.Println("Error saving expenses:", err)
		os.Exit(1)
	}
}
