package main

import (
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	Job    string
	Salary float64
}

// ReadData method to read data into the struct from user input
func (p Person) ReadData() Person {
	fmt.Print("Enter Name: ")
	fmt.Scan(&p.Name)

	fmt.Print("Enter Age: ")
	fmt.Scan(&p.Age)

	fmt.Print("Enter Job: ")
	fmt.Scan(&p.Job)

	fmt.Print("Enter Salary: ")
	fmt.Scan(&p.Salary)

	return p
}

// PrintData method to print all fields in a formatted manner
func (p Person) PrintData() {
	fmt.Println("================================")
	fmt.Printf("Name:   %s\n", p.Name)
	fmt.Printf("Age:    %d years\n", p.Age)
	fmt.Printf("Job:    %s\n", p.Job)
	fmt.Printf("Salary: $%.2f\n", p.Salary)
	fmt.Println("================================")
}

func main() {
	fmt.Println("Person Struct with Associated Methods - LAB 3")
	fmt.Println("=============================================")

	// Create first Person object
	fmt.Println("\n--- Creating First Person ---")
	var person1 Person
	person1 = person1.ReadData()

	// Create second Person object
	fmt.Println("\n--- Creating Second Person ---")
	var person2 Person
	person2 = person2.ReadData()

	// Display both persons' data
	fmt.Println("\n--- Person 1 Details ---")
	person1.PrintData()

	fmt.Println("\n--- Person 2 Details ---")
	person2.PrintData()

	// Example with pre-filled data for demonstration
	fmt.Println("\n--- Example Person (Pre-filled) ---")
	person3 := Person{
		Name:   "John Doe",
		Age:    30,
		Job:    "Software Engineer",
		Salary: 75000.50,
	}
	person3.PrintData()
}