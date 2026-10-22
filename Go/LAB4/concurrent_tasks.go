package main

import (
	"fmt"
	"sync"
	"time"
)

// Result struct to hold task results with identification
type Result struct {
	TaskName string
	Input    int
	Output   interface{}
}

// Goroutine 1: Calculate squares of numbers
func calculateSquares(numbers []int, ch chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[GOROUTINE 1] 🟢 STARTED: Calculating Squares")
	
	for _, num := range numbers {
		square := num * num
		time.Sleep(100 * time.Millisecond) // Simulate computation time
		ch <- Result{
			TaskName: "Square",
			Input:    num,
			Output:   square,
		}
		fmt.Printf("[GOROUTINE 1] Computed: %d² = %d\n", num, square)
	}
	
	fmt.Println("[GOROUTINE 1] 🔴 COMPLETED: All squares calculated")
}

// Goroutine 2: Calculate cubes of numbers
func calculateCubes(numbers []int, ch chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[GOROUTINE 2] 🟢 STARTED: Calculating Cubes")
	
	for _, num := range numbers {
		cube := num * num * num
		time.Sleep(150 * time.Millisecond) // Simulate different computation time
		ch <- Result{
			TaskName: "Cube",
			Input:    num,
			Output:   cube,
		}
		fmt.Printf("[GOROUTINE 2] Computed: %d³ = %d\n", num, cube)
	}
	
	fmt.Println("[GOROUTINE 2] 🔴 COMPLETED: All cubes calculated")
}

// Goroutine 3: Calculate Fibonacci numbers
func calculateFibonacci(count int, ch chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[GOROUTINE 3] 🟢 STARTED: Calculating Fibonacci Numbers")
	
	for i := 0; i < count; i++ {
		fib := fibonacci(i)
		time.Sleep(80 * time.Millisecond) // Simulate computation time
		ch <- Result{
			TaskName: "Fibonacci",
			Input:    i,
			Output:   fib,
		}
		fmt.Printf("[GOROUTINE 3] Computed: Fib(%d) = %d\n", i, fib)
	}
	
	fmt.Println("[GOROUTINE 3] 🔴 COMPLETED: All Fibonacci numbers calculated")
}

// Goroutine 4: Calculate factorials (additional task)
func calculateFactorials(numbers []int, ch chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[GOROUTINE 4] 🟢 STARTED: Calculating Factorials")
	
	for _, num := range numbers {
		factorial := calculateFactorial(num)
		time.Sleep(120 * time.Millisecond) // Simulate computation time
		ch <- Result{
			TaskName: "Factorial",
			Input:    num,
			Output:   factorial,
		}
		fmt.Printf("[GOROUTINE 4] Computed: %d! = %d\n", num, factorial)
	}
	
	fmt.Println("[GOROUTINE 4] 🔴 COMPLETED: All factorials calculated")
}

// Helper function to calculate Fibonacci number
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Helper function to calculate factorial
func calculateFactorial(n int) int {
	if n <= 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║   LAB 4: Concurrent Tasks Using Goroutines and Channels    ║")
	fmt.Println("║                    (Experiment 6)                          ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
	fmt.Println()

	// Create a channel to collect results
	resultChannel := make(chan Result, 20) // Buffered channel for better performance
	
	// WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup

	// Input data for computations
	numbers := []int{2, 3, 4, 5}
	fibCount := 8

	fmt.Println("🚀 MAIN: Launching concurrent goroutines...")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println()

	// Launch goroutines concurrently
	wg.Add(4)
	
	go calculateSquares(numbers, resultChannel, &wg)
	go calculateCubes(numbers, resultChannel, &wg)
	go calculateFibonacci(fibCount, resultChannel, &wg)
	go calculateFactorials(numbers, resultChannel, &wg)

	// Close channel after all goroutines complete
	go func() {
		wg.Wait()
		close(resultChannel)
		fmt.Println()
		fmt.Println("═══════════════════════════════════════════════════════════")
		fmt.Println("✅ MAIN: All goroutines completed")
	}()

	// Collect results from channel
	fmt.Println()
	fmt.Println("📊 MAIN: Collecting results from channel...")
	fmt.Println("───────────────────────────────────────────────────────────")
	
	// Storage for organized results
	squares := []Result{}
	cubes := []Result{}
	fibs := []Result{}
	factorials := []Result{}

	// Receive all results from channel
	for result := range resultChannel {
		switch result.TaskName {
		case "Square":
			squares = append(squares, result)
		case "Cube":
			cubes = append(cubes, result)
		case "Fibonacci":
			fibs = append(fibs, result)
		case "Factorial":
			factorials = append(factorials, result)
		}
	}

	// Print organized results
	fmt.Println()
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║                     FINAL RESULTS                          ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
	fmt.Println()

	// Print Squares
	fmt.Println("📐 SQUARES:")
	for _, r := range squares {
		fmt.Printf("   %d² = %v\n", r.Input, r.Output)
	}
	fmt.Println()

	// Print Cubes
	fmt.Println("📦 CUBES:")
	for _, r := range cubes {
		fmt.Printf("   %d³ = %v\n", r.Input, r.Output)
	}
	fmt.Println()

	// Print Fibonacci Numbers
	fmt.Println("🔢 FIBONACCI SEQUENCE:")
	fmt.Print("   ")
	for i, r := range fibs {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("F(%d)=%v", r.Input, r.Output)
	}
	fmt.Println()
	fmt.Println()

	// Print Factorials
	fmt.Println("🎯 FACTORIALS:")
	for _, r := range factorials {
		fmt.Printf("   %d! = %v\n", r.Input, r.Output)
	}
	fmt.Println()

	// Summary
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Printf("📈 Total results collected: %d\n", len(squares)+len(cubes)+len(fibs)+len(factorials))
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println()
	fmt.Println("✨ Program completed successfully!")
	fmt.Println("🔍 Notice: Goroutines executed concurrently as shown by")
	fmt.Println("   the interleaved start/completion messages above.")
}