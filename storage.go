package main

import (
	"os"
	"encoding/json"
)

type Storage[T any] struct {
	FileName string
}

// NewStorage creates a new Storage instance with the given fileName
func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{FileName: fileName}
}

func (s *Storage[T]) Save(data T) error {
	
	// Implementation for saving data to the storage file
	// Replace this with actual file writing logic

	fileData, err := json.MarshalIndent(data, "", "   ")

	if err!= nil {
        return err
    }

	return os.WriteFile(s.FileName, fileData, 0644)

}

func (s *Storage[T]) Load(data *T) error {
	
	fileData, err := os.ReadFile(s.FileName)

	if err!= nil {
        return err
    }

	return json.Unmarshal(fileData, data)
}