package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type CommandFlags struct {
	Add  string
	List bool
}

func newCommandFlags() *CommandFlags {
	cf := CommandFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add an expense in format: 'title,budget,spent'")
	flag.BoolVar(&cf.List, "list", false, "List all expenses")

	flag.Parse()
	return &cf
}

func (cf *CommandFlags) Execute(expenses *Expenses) error {
	switch {
	case cf.List:
		expenses.PrintAll()
		return nil
	case cf.Add != "":
		title, budget, spent, err := parseAddCommand(cf.Add)
		if err != nil {
			return fmt.Errorf("failed to add expense: %v", err)
		}
		expenses.Add(title, budget, spent)
		return nil
	default:
		flag.Usage()
		return nil
	}
}

func parseAddCommand(addCommand string) (string, int, int, error) {
	parts := strings.Split(addCommand, ",")
	if len(parts) != 3 {
		return "", 0, 0, fmt.Errorf("invalid format: please use 'title,budget,spent'")
	}

	title := strings.TrimSpace(parts[0])
	if title == "" {
		return "", 0, 0, fmt.Errorf("title cannot be empty")
	}

	budget, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return "", 0, 0, fmt.Errorf("invalid budget value: %v", err)
	}

	spent, err := strconv.Atoi(strings.TrimSpace(parts[2]))
	if err != nil {
		return "", 0, 0, fmt.Errorf("invalid spent value: %v", err)
	}

	if budget < 0 {
		return "", 0, 0, fmt.Errorf("budget cannot be negative")
	}

	if spent < 0 {
		return "", 0, 0, fmt.Errorf("spent amount cannot be negative")
	}

	return title, budget, spent, nil
}
