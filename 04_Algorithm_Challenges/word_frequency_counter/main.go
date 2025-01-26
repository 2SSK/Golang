package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func writeFile() {
	reader := bufio.NewReader(os.Stdin)
	var content string
	fmt.Println("Enter the content to write in the file")

	for {
		// Read a line of input
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input", err)
			return
		}

		// If the line is empty, stop reading
		if line == "\n" {
			break
		}

		// Append the line to the content
		content += line
	}

	// Open the file
	file, err := os.Create("input.txt")
	if err != nil {
		fmt.Println("Error creating file", err)
		return
	}
	defer file.Close()

	// Write the content to the writeFile
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file", err)
		return
	}
}

func main() {
	writeFile()

	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	defer file.Close()

	// Create a map to store the frequency of each character
	mp := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, char := range line {
			if char != ' ' && unicode.IsLetter(char) {
				mp[strings.ToLower(string(char))]++
			}
		}
	}

	// check for errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	// sort the map
	var keys []string
	for key := range mp {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Print the frequency of each character
	for _, key := range keys {
		fmt.Printf("%v : %v\n", key, mp[key])
	}

	fmt.Println()
}
