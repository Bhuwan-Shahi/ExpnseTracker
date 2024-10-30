package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	fileName string
}

func newStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{fileName: fileName}
}

func (s *Storage[T]) Save(data T) error {
	// Pretty print JSON for better readability
	fileData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.fileName, fileData, 0644)
}

func (s *Storage[T]) Load(data *T) error {
	// Check if file exists
	if _, err := os.Stat(s.fileName); os.IsNotExist(err) {
		return err
	}

	fileData, err := os.ReadFile(s.fileName)
	if err != nil {
		return err
	}

	// Handle empty file
	if len(fileData) == 0 {
		return os.ErrNotExist
	}

	return json.Unmarshal(fileData, data)
}
