# LAB 3 - Struct Person with Associated Methods (Experiment 4)

## Objective
Define a Go struct named `Person` with fields and associated methods to demonstrate struct usage and method implementation in Go.

## Requirements Implemented

### 1. Person Struct Definition
- **Name** (string): Person's name
- **Age** (int): Person's age in years
- **Job** (string): Person's job title/profession
- **Salary** (float64): Person's salary

### 2. Associated Methods

#### ReadData() Method
- **Receiver**: Value receiver `(p Person)` - uses value receiver as required
- **Return Type**: Returns updated `Person` struct
- **Purpose**: Reads user input for all struct fields and returns modified struct
- **Features**:
  - Interactive input prompting
  - Input validation for Age and Salary fields
  - Error handling with default values
  - Returns the modified Person struct

#### PrintData() Method
- **Receiver**: Value receiver `(p Person)` - read-only access
- **Purpose**: Displays all struct fields in a formatted manner
- **Features**:
  - Professional formatting with borders
  - Currency formatting for salary ($XX.XX)
  - Clear field labels

### 3. Main Function Implementation
- Creates two Person objects using user input
- Uses value receivers for both methods
- Assigns returned Person struct from ReadData() method
- Calls both methods on each Person object
- Includes a pre-filled example for demonstration

## How to Run

```bash
# Navigate to LAB3 directory
cd LAB3

# Run the program
go run main.go

# Or compile and run
go build -o person.exe main.go
./person.exe
```

## Sample Output
```
Person Struct with Associated Methods - LAB 3
=============================================

--- Creating First Person ---
Enter Name: Alice Johnson
Enter Age: 28
Enter Job: Data Scientist
Enter Salary: 85000

--- Creating Second Person ---
Enter Name: Bob Smith
Enter Age: 35
Enter Job: Project Manager
Enter Salary: 95000

--- Person 1 Details ---
================================
Name:   Alice Johnson
Age:    28 years
Job:    Data Scientist
Salary: $85000.00
================================

--- Person 2 Details ---
================================
Name:   Bob Smith
Age:    35 years
Job:    Project Manager
Salary: $95000.00
================================
```

## Key Go Concepts Demonstrated
- Struct definition and field access
- Value receivers for all methods (no pointers)
- Method implementation for structs that return modified struct
- User input handling with bufio
- String parsing and type conversion
- Error handling for invalid input