package main

import (
	"fmt"
)

// Person struct for demonstrating pointer allocation with new()
type Person struct {
	Name string
	Age  int
	City string
}

// Function that accepts a pointer parameter and modifies the original variable
func modifyValue(ptr *int) {
	fmt.Printf("  Inside function - Address: %p, Value: %d\n", ptr, *ptr)
	*ptr = *ptr * 2 // Double the value
	fmt.Printf("  After modification - Address: %p, Value: %d\n", ptr, *ptr)
}

// Function to modify struct fields through pointer
func updatePerson(p *Person, newName string, newAge int, newCity string) {
	fmt.Printf("  Before update - Name: %s, Age: %d, City: %s\n", p.Name, p.Age, p.City)
	p.Name = newName
	p.Age = newAge
	p.City = newCity
	fmt.Printf("  After update - Name: %s, Age: %d, City: %s\n", p.Name, p.Age, p.City)
}

func main() {
	fmt.Println("=== LAB 3: Pointer Manipulation ===\n")

	// (1) Basic pointer operations - declare variable, print address and access value
	fmt.Println("1. Basic Pointer Operations:")
	fmt.Println("   Declaring variable and demonstrating & and * operators")
	
	var number int = 42
	var ptr *int = &number
	
	fmt.Printf("   Variable 'number' value: %d\n", number)
	fmt.Printf("   Variable 'number' address using &: %p\n", &number)
	fmt.Printf("   Pointer 'ptr' stores address: %p\n", ptr)
	fmt.Printf("   Value at address using *ptr: %d\n", *ptr)
	fmt.Printf("   Address of pointer itself: %p\n", &ptr)
	
	// Demonstrate pointer arithmetic concept
	fmt.Printf("   Changing value through pointer...\n")
	*ptr = 100
	fmt.Printf("   New value of 'number': %d\n", number)
	fmt.Printf("   New value through *ptr: %d\n", *ptr)
	
	fmt.Println()

	// (2) Pass-by-reference using pointer parameter
	fmt.Println("2. Pass-by-Reference Function:")
	fmt.Println("   Function that modifies original variable through pointer")
	
	var originalValue int = 25
	fmt.Printf("   Before function call - Value: %d, Address: %p\n", originalValue, &originalValue)
	
	fmt.Println("   Calling modifyValue function...")
	modifyValue(&originalValue)
	
	fmt.Printf("   After function call - Value: %d, Address: %p\n", originalValue, &originalValue)
	fmt.Println()

	// (3) Struct allocation using new() and pointer access
	fmt.Println("3. Struct Allocation with new() and Pointer Access:")
	fmt.Println("   Allocating Person struct using new() and accessing fields through pointer")
	
	// Allocate struct using new()
	personPtr := new(Person)
	fmt.Printf("   Allocated Person struct at address: %p\n", personPtr)
	fmt.Printf("   Initial values - Name: '%s', Age: %d, City: '%s'\n", 
		personPtr.Name, personPtr.Age, personPtr.City)
	
	// Access and modify fields through pointer
	personPtr.Name = "John Doe"
	personPtr.Age = 30
	personPtr.City = "New York"
	
	fmt.Printf("   After direct assignment - Name: '%s', Age: %d, City: '%s'\n", 
		personPtr.Name, personPtr.Age, personPtr.City)
	
	// Alternative syntax using (*pointer).field
	(*personPtr).Name = "Jane Smith"
	(*personPtr).Age = 25
	(*personPtr).City = "San Francisco"
	
	fmt.Printf("   Using (*ptr).field syntax - Name: '%s', Age: %d, City: '%s'\n", 
		(*personPtr).Name, (*personPtr).Age, (*personPtr).City)
	
	// Using function to modify through pointer
	fmt.Println("   Using function to modify struct through pointer:")
	updatePerson(personPtr, "Alice Johnson", 28, "Chicago")
	
	fmt.Printf("   Final values - Name: '%s', Age: %d, City: '%s'\n", 
		personPtr.Name, personPtr.Age, personPtr.City)
	
	fmt.Println()

	// Additional demonstration: Multiple pointers to same variable
	fmt.Println("4. Additional: Multiple Pointers to Same Variable")
	var sharedVar int = 77
	ptr1 := &sharedVar
	ptr2 := &sharedVar
	
	fmt.Printf("   Original value: %d\n", sharedVar)
	fmt.Printf("   ptr1 points to: %p, value: %d\n", ptr1, *ptr1)
	fmt.Printf("   ptr2 points to: %p, value: %d\n", ptr2, *ptr2)
	
	*ptr1 = 99
	fmt.Printf("   After changing *ptr1 to 99:\n")
	fmt.Printf("   sharedVar: %d\n", sharedVar)
	fmt.Printf("   *ptr2: %d\n", *ptr2)
	
	fmt.Println("\n=== End of Pointer Manipulation Demo ===")
}