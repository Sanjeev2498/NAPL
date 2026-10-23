package main

import (
	"errors"
	"fmt"
)

// Item represents a product in inventory
type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

// Inventory manages the collection of items
type Inventory struct {
	items []Item
}

// NewInventory creates a new inventory
func NewInventory() *Inventory {
	return &Inventory{
		items: make([]Item, 0),
	}
}

// AddItem adds a new item to inventory
func (inv *Inventory) AddItem(id int, name string, quantity int, price float64) error {
	// Validation
	if id <= 0 {
		return errors.New("ID must be positive")
	}
	if name == "" {
		return errors.New("name cannot be empty")
	}
	if quantity < 0 {
		return errors.New("quantity cannot be negative")
	}
	if price < 0 {
		return errors.New("price cannot be negative")
	}

	// Check if ID already exists
	for _, item := range inv.items {
		if item.ID == id {
			return errors.New("item with this ID already exists")
		}
	}

	// Add item
	newItem := Item{
		ID:       id,
		Name:     name,
		Quantity: quantity,
		Price:    price,
	}
	inv.items = append(inv.items, newItem)
	return nil
}

// UpdateItem updates an existing item
func (inv *Inventory) UpdateItem(id int, name string, quantity int, price float64) error {
	// Validation
	if name == "" {
		return errors.New("name cannot be empty")
	}
	if quantity < 0 {
		return errors.New("quantity cannot be negative")
	}
	if price < 0 {
		return errors.New("price cannot be negative")
	}

	// Find and update item
	for i, item := range inv.items {
		if item.ID == id {
			inv.items[i].Name = name
			inv.items[i].Quantity = quantity
			inv.items[i].Price = price
			return nil
		}
	}

	return errors.New("item not found")
}

// RemoveItem removes an item from inventory
func (inv *Inventory) RemoveItem(id int) error {
	for i, item := range inv.items {
		if item.ID == id {
			// Remove item by slicing
			inv.items = append(inv.items[:i], inv.items[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}

// GetItem retrieves an item by ID
func (inv *Inventory) GetItem(id int) (Item, error) {
	for _, item := range inv.items {
		if item.ID == id {
			return item, nil
		}
	}
	return Item{}, errors.New("item not found")
}

// GetAllItems returns all items
func (inv *Inventory) GetAllItems() []Item {
	return inv.items
}

// Statistics holds summary data
type Statistics struct {
	TotalItems int
	TotalValue float64
}

// CalculateStats computes statistics using goroutine and channel
func (inv *Inventory) CalculateStats(result chan<- Statistics) {
	stats := Statistics{
		TotalItems: len(inv.items),
		TotalValue: 0,
	}

	for _, item := range inv.items {
		stats.TotalValue += float64(item.Quantity) * item.Price
	}

	result <- stats
}

// DisplayItems prints all items in a formatted way
func (inv *Inventory) DisplayItems() {
	if len(inv.items) == 0 {
		fmt.Println("No items in inventory")
		return
	}

	fmt.Println("\n--- Inventory Items ---")
	fmt.Printf("%-5s %-20s %-10s %-10s %-10s\n", "ID", "Name", "Quantity", "Price", "Total")
	fmt.Println("----------------------------------------------------------------")

	for _, item := range inv.items {
		total := float64(item.Quantity) * item.Price
		fmt.Printf("%-5d %-20s %-10d $%-9.2f $%-9.2f\n",
			item.ID, item.Name, item.Quantity, item.Price, total)
	}
	fmt.Println()
}
