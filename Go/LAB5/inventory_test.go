package main

import (
	"testing"
)

// Test adding items
func TestAddItem(t *testing.T) {
	inv := NewInventory()

	// Test valid item
	err := inv.AddItem(1, "Laptop", 10, 999.99)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Check if item was added
	if len(inv.items) != 1 {
		t.Errorf("Expected 1 item, got: %d", len(inv.items))
	}

	// Test duplicate ID
	err = inv.AddItem(1, "Mouse", 5, 25.50)
	if err == nil {
		t.Error("Expected error for duplicate ID, got none")
	}

	// Test invalid quantity
	err = inv.AddItem(2, "Keyboard", -5, 50.00)
	if err == nil {
		t.Error("Expected error for negative quantity, got none")
	}

	// Test empty name
	err = inv.AddItem(3, "", 10, 50.00)
	if err == nil {
		t.Error("Expected error for empty name, got none")
	}
}

// Test updating items
func TestUpdateItem(t *testing.T) {
	inv := NewInventory()
	inv.AddItem(1, "Laptop", 10, 999.99)

	// Test valid update
	err := inv.UpdateItem(1, "Gaming Laptop", 15, 1299.99)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Check if item was updated
	item, _ := inv.GetItem(1)
	if item.Name != "Gaming Laptop" {
		t.Errorf("Expected name 'Gaming Laptop', got: %s", item.Name)
	}
	if item.Quantity != 15 {
		t.Errorf("Expected quantity 15, got: %d", item.Quantity)
	}

	// Test updating non-existent item
	err = inv.UpdateItem(999, "Test", 5, 10.00)
	if err == nil {
		t.Error("Expected error for non-existent item, got none")
	}
}

// Test removing items
func TestRemoveItem(t *testing.T) {
	inv := NewInventory()
	inv.AddItem(1, "Laptop", 10, 999.99)
	inv.AddItem(2, "Mouse", 50, 25.50)

	// Test valid removal
	err := inv.RemoveItem(1)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Check if item was removed
	if len(inv.items) != 1 {
		t.Errorf("Expected 1 item, got: %d", len(inv.items))
	}

	// Test removing non-existent item
	err = inv.RemoveItem(999)
	if err == nil {
		t.Error("Expected error for non-existent item, got none")
	}
}

// Test calculating statistics
func TestCalculateStats(t *testing.T) {
	inv := NewInventory()
	inv.AddItem(1, "Laptop", 10, 100.00)
	inv.AddItem(2, "Mouse", 20, 50.00)

	statsChan := make(chan Statistics)
	go inv.CalculateStats(statsChan)
	stats := <-statsChan

	// Expected: 10*100 + 20*50 = 1000 + 1000 = 2000
	expectedValue := 2000.0
	if stats.TotalValue != expectedValue {
		t.Errorf("Expected total value %.2f, got: %.2f", expectedValue, stats.TotalValue)
	}

	if stats.TotalItems != 2 {
		t.Errorf("Expected 2 items, got: %d", stats.TotalItems)
	}
}

// Test getting item by ID
func TestGetItem(t *testing.T) {
	inv := NewInventory()
	inv.AddItem(1, "Laptop", 10, 999.99)

	// Test valid get
	item, err := inv.GetItem(1)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if item.Name != "Laptop" {
		t.Errorf("Expected name 'Laptop', got: %s", item.Name)
	}

	// Test getting non-existent item
	_, err = inv.GetItem(999)
	if err == nil {
		t.Error("Expected error for non-existent item, got none")
	}
}
