package main

import (
	"fmt"
)

func main() {
	var a, b float64

	// Input for first number
	for {
		fmt.Print("Enter first number (-1000 to 1000): ")

		_, err := fmt.Scan(&a)

		if err != nil {
			fmt.Println("Invalid input! Please enter a numeric value.")

			// Clear invalid input
			var discard string
			fmt.Scan(&discard)
			continue
		}

		if a < -1000 || a > 1000 {
			fmt.Println("Number out of range! Enter a value between -1000 and 1000.")
			continue
		}

		break
	}

	// Input for second number
	for {
		fmt.Print("Enter second number (-1000 to 1000): ")

		_, err := fmt.Scan(&b)

		if err != nil {
			fmt.Println("Invalid input! Please enter a numeric value.")

			// Clear invalid input
			var discard string
			fmt.Scan(&discard)
			continue
		}

		if b < -1000 || b > 1000 {
			fmt.Println("Number out of range! Enter a value between -1000 and 1000.")
			continue
		}

		break
	}

	// Perform calculations
	addition := a + b
	subtraction := a - b
	multiplication := a * b

	fmt.Println("\n--- Calculator Results ---")
	fmt.Printf("Addition       = %.2f\n", addition)
	fmt.Printf("Subtraction    = %.2f\n", subtraction)
	fmt.Printf("Multiplication = %.2f\n", multiplication)
}
