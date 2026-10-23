# LAB 5: Inventory Management System

## Overview
A simple inventory management module that uses core Go concepts including structs, slices, error handling, goroutines, channels, and unit testing.

## Features

### 1. Item Struct
- ID (int)
- Name (string)
- Quantity (int)
- Price (float64)

### 2. Operations
- **Add Item**: Add new items with validation
- **Update Item**: Modify existing items
- **Remove Item**: Delete items from inventory
- **Get Item**: Retrieve item by ID
- **Display Items**: Show all items in formatted table

### 3. Validation & Error Handling
- ID must be positive and unique
- Name cannot be empty
- Quantity and price cannot be negative
- Proper error messages for all operations

### 4. Concurrent Statistics
- Uses goroutine and channel to calculate:
  - Total number of items
  - Total inventory value

### 5. Unit Tests
Five test functions covering:
- Adding items (valid and invalid cases)
- Updating items
- Removing items
- Calculating statistics
- Getting items by ID

## How to Run

### Run the program
```bash
cd Go/LAB5
go run .
```

### Run tests
```bash
go test
```

### Run tests with details
```bash
go test -v
```

## Files
- `inventory.go` - Core inventory logic
- `main.go` - Main program demonstrating functionality
- `inventory_test.go` - Unit tests
- `go.mod` - Module configuration
- `README.md` - Documentation

## Sample Output
```
=== Inventory Management System ===

Adding items to inventory...
Items added successfully!

--- Inventory Items ---
ID    Name                 Quantity   Price      Total     
----------------------------------------------------------------
1     Laptop               10         $999.99    $9999.90  
2     Mouse                50         $25.50     $1275.00  
3     Keyboard             30         $75.00     $2250.00  
4     Monitor              15         $299.99    $4499.85  

--- Inventory Summary ---
Total Items: 3
Total Value: $13524.90
```

## Concepts Demonstrated
✅ Structs and methods  
✅ Slices for data storage  
✅ Error handling with validation  
✅ Goroutines and channels  
✅ Unit testing with table-driven tests  
✅ Input validation  
✅ Clean, maintainable code structure
