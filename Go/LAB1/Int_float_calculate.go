package main

import "fmt"

func main() {
	// Integer calculation
	var int1, int2 int

	fmt.Println("----- Integer Calculator -----")

	fmt.Print("Enter first integer: ")
	fmt.Scan(&int1)

	fmt.Print("Enter second integer: ")
	fmt.Scan(&int2)

	fmt.Println("Addition       =", int1+int2)
	fmt.Println("Subtraction    =", int1-int2)
	fmt.Println("Multiplication =", int1*int2)

	// Floating-point calculation
	var float1, float2 float64

	fmt.Println("\n----- Floating-Point Calculator -----")

	fmt.Print("Enter first decimal number: ")
	fmt.Scan(&float1)

	fmt.Print("Enter second decimal number: ")
	fmt.Scan(&float2)

	fmt.Printf("Addition       = %.2f\n", float1+float2)
	fmt.Printf("Subtraction    = %.2f\n", float1-float2)
	fmt.Printf("Multiplication = %.2f\n", float1*float2)
}