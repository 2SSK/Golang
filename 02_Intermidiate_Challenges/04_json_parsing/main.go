package main

import "fmt"

type Fullname struct {
	FirstName string
	LastName  string
}

type Employee struct {
	Name  Fullname
	Age   int
	Email string
}

func printEmployee(emp Employee) {
	fmt.Println("Name:", emp.Name.FirstName, emp.Name.LastName)
	fmt.Println("Age:", emp.Age)
	fmt.Println("Email:", emp.Email)
	fmt.Println()
}

func main() {
	// Create a slice of Employee
	var employees []Employee

	err := ParseJsonFile("data.json", &employees)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the slice of Employee
	for _, e := range employees {
		printEmployee(e)
	}
}
