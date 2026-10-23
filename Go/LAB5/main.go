package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Inventory Management System ===")
	fmt.Println()

	// Create inventory
	inventory := NewInventory()

	// Add items
	fmt.Println("Adding items to inventory...")
	inventory.AddItem(1, "Laptop", 10, 999.99)
	inventory.AddItem(2, "Mouse", 50, 25.50)
	inventory.AddItem(3, "Keyboard", 30, 75.00)
	inventory.AddItem(4, "Monitor", 15, 299.99)
	fmt.Println("Items added successfully!")

	// Display items
	inventory.DisplayItems()

	// Update an item
	fmt.Println("Updating Mouse quantity to 45...")
	err := inventory.UpdateItem(2, "Mouse", 45, 25.50)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Item updated successfully!")
	}
	inventory.DisplayItems()

	// Remove an item
	fmt.Println("Removing Monitor (ID: 4)...")
	err = inventory.RemoveItem(4)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Item removed successfully!")
	}
	inventory.DisplayItems()

	// Calculate statistics using goroutine
	fmt.Println("Calculating inventory statistics...")
	statsChan := make(chan Statistics)

	go inventory.CalculateStats(statsChan)

	stats := <-statsChan
	close(statsChan)

	fmt.Printf("\n--- Inventory Summary ---\n")
	fmt.Printf("Total Items: %d\n", stats.TotalItems)
	fmt.Printf("Total Value: $%.2f\n", stats.TotalValue)
	fmt.Println("\nProgram completed successfully!")
}
