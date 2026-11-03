package main

import (
	"fmt"
	"math"
	"strings"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func squareRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("cannot find the square root of a negative number")
	}
	return math.Sqrt(number), nil
}

func greet(name string) (string, error) {
	if strings.TrimSpace(name) == "" {
		return "", fmt.Errorf("name cannot be empty")
	}
	return "Hello, " + name + "!", nil
}

func runErrorHandlingExamples() {
	fmt.Println()
	fmt.Println("LAB 4: Error Handling Examples (Experiment 7)")

	if result, err := divide(10, 2); err != nil {
		fmt.Println("Division error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	if result, err := divide(10, 0); err != nil {
		fmt.Println("Division error:", err)
	} else {
		fmt.Println("10 / 0 =", result)
	}

	if result, err := squareRoot(25); err != nil {
		fmt.Println("Square root error:", err)
	} else {
		fmt.Println("Square root of 25 =", result)
	}

	if result, err := squareRoot(-4); err != nil {
		fmt.Println("Square root error:", err)
	} else {
		fmt.Println("Square root of -4 =", result)
	}

	if message, err := greet("Sanjeev"); err != nil {
		fmt.Println("Input error:", err)
	} else {
		fmt.Println(message)
	}

	if message, err := greet(""); err != nil {
		fmt.Println("Input error:", err)
	} else {
		fmt.Println(message)
	}
}
