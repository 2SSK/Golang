package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func ParseJsonFile(filename string, data interface{}) error {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Error opening JSON file: %v", err)
	}
	defer file.Close()

	// Decode the JSON file
	decoder := json.NewDecoder(file)
	err = decoder.Decode(data)
	if err != nil {
		return fmt.Errorf("Error decoding JSON file: %v", err)
	}

	return nil
}
