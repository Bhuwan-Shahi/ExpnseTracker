package main

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/aquasecurity/table"
)

type Expense struct {
	Title    string `json:"title"`
	Budget   int    `json:"budget"`
	Spent    int    `json:"spent"`
	Profit   int    `json:"profit"`
	IsProfit bool   `json:"isProfit"`
}

// Ensure Expenses type is properly initialized
type Expenses []Expense

// MarshalJSON ensures the JSON is always an array
func (e Expenses) MarshalJSON() ([]byte, error) {
	if e == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]Expense(e))
}

// UnmarshalJSON ensures we can handle both null and array inputs
func (e *Expenses) UnmarshalJSON(data []byte) error {
	// Handle null or empty array
	if string(data) == "null" || string(data) == "[]" {
		*e = make(Expenses, 0)
		return nil
	}

	// Normal array unmarshal
	var expenses []Expense
	if err := json.Unmarshal(data, &expenses); err != nil {
		return err
	}
	*e = expenses
	return nil
}

// Add adds a new expense to the expenses slice
func (e *Expenses) Add(title string, budget, spent int) {
	if *e == nil {
		*e = make(Expenses, 0)
	}

	profit := budget - spent
	expense := Expense{
		Title:    title,
		Budget:   budget,
		Spent:    spent,
		Profit:   profit,
		IsProfit: profit > 0,
	}

	*e = append(*e, expense)
}

// UpdateProfitStatus updates the IsProfit status for all expenses
func (e *Expenses) UpdateProfitStatus() {
	for i := range *e {
		(*e)[i].IsProfit = (*e)[i].Profit > 0
	}
}

// PrintAll prints all expenses in a formatted table
func (e *Expenses) PrintAll() {
	t := table.New(os.Stdout)
	t.SetRowLines(false)
	t.SetHeaders("#", "Title", "Budget", "Spent", "Profit", "IsProfit")

	for i, exp := range *e {
		var isProfit string
		if exp.IsProfit {
			isProfit = "✅"
		} else {
			isProfit = "❌"
		}

		t.AddRow(
			strconv.Itoa(i),
			exp.Title,
			strconv.Itoa(exp.Budget),
			strconv.Itoa(exp.Spent),
			strconv.Itoa(exp.Profit),
			isProfit,
		)
	}

	t.Render()
}
