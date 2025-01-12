package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Function to create file
func createFile(fileName string) {
	fmt.Printf("\n\tCreating file with name \"%v\".\n", fileName)

	// Create a file
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Printf("Error while creating file : %v", err)
		return
	}

	fmt.Printf("File \"%v\" has been created successfully!\n", file.Name())
	defer file.Close()
}

// Function to write in a file
func writeFile(file string, data string) {
	fmt.Println("\nWriting in file: ", file)

	// Open the file in append mode if it exists
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error while accessing file(%v): %v\n", f, err)
		return
	}
	defer f.Close()

	// Check if the file ends with a newline
	fileInfo, err := f.Stat()
	if err == nil && fileInfo.Size() > 0 {
		_, err = f.WriteString("\n")
		if err != nil {
			fmt.Printf("Error while adding new line : %v\n", err)
		}
	}

	// Write the data in file
	_, err = f.WriteString(data)
	if err != nil {
		fmt.Printf("Error while writing file : %v\n", err)
		return
	}
	fmt.Printf("\nFile has been modified successfully!!\n")

}

// Function to read in a file
func readFile(file string) {
	fmt.Printf("\nReading file \"%v\".\n", file)
	fmt.Println("---------------------------------------")
	// Access the file
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("\nError access file \"%v\" : %v\n", file, err)
		return
	}
	defer f.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(f)
	i := 1 // index of lines
	for scanner.Scan() {
		fmt.Printf("%v. %v\n", i, scanner.Text())
		i++
	}

	// Check for error
	if err := scanner.Err(); err != nil {
		fmt.Printf("\nError while scanning file '%v' : %v", file, err)
		return
	}
}

func takeFilename() string {
	// Take filename from user
	var file string
	fmt.Print("\nEnter the file name : ")
	fmt.Scanln(&file)
	file = strings.TrimSpace(file)

	return string(file)
}

func takeData() string {
	// Take data from user
	fmt.Printf("\n\tWriting in the file\nEnter the data(end with ctrl + d) : \n")
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("Error while taking data: %v", err)
		return ""
	}
	return string(data)
}

// Check if file exists and run the function
func checkFile() (string, bool) {
	file := takeFilename()

	// Check if file already exist
	_, err := os.Stat(file)
	if err != nil {
		return file, false
	}
	return file, true
}

// Run the function based on the command
func superCmd(cmd string) {
	file, fileExists := checkFile()

	// Run the function based on the command
	if cmd == "create" {
		if fileExists {
			fmt.Println("\nFile already exists")
			return
		} else {
			createFile(file)
		}
	} else if cmd == "write" {
		if fileExists {
			writeFile(file, takeData())
		} else {
			fmt.Println("\nFile does not exist")
		}
	} else if cmd == "read" {
		if fileExists {
			readFile(file)
		} else {
			fmt.Println("\nFile does not exist")
		}
	}
}

// Display options
func options() {
	fmt.Println("\n---------------------------------------")
	fmt.Println("\tFile Handling in GoLang")
	fmt.Println("---------------------------------------")
	fmt.Printf("\n1. Create a file\n2. Write in a file\n3. Read a file\n4. Exit\n")
}

// Ui for the user
func starterUi() {

	choice := 0

	options()

	for {
		exit := false
		// Display the options
		fmt.Print("\nEnter your choice (1,2,3,4) : ")

		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Printf("Invalid choice, please enter a number")
			return
		}

		switch choice {
		case 1:
			superCmd("create")
		case 2:
			superCmd("write")
		case 3:
			superCmd("read")
		case 4:
			fmt.Println("Exiting the program...")
			exit = true
			break
		default:
			fmt.Printf("Invalid choice\n\n")
		}

		if exit {
			break
		}
	}
}

// Main function
func main() {
	starterUi()
}
