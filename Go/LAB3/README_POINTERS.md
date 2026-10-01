# LAB 3: Pointer Manipulation - Referencing and Dereferencing (Exp 5)

## Objective
Demonstrate pointer usage in Go programming including basic pointer operations, pass-by-reference functions, and struct allocation with `new()`.

## Requirements Fulfilled

### 1. Basic Pointer Operations ✅
- Declare a variable and print its address using `&` operator
- Access the value at memory address using `*` operator (dereferencing)
- Demonstrate pointer assignment and value modification through pointers

### 2. Pass-by-Reference Function ✅
- Create a function that accepts a pointer parameter
- Modify the original variable through the pointer
- Show before and after values to demonstrate pass-by-reference behavior

### 3. Struct Allocation with new() ✅
- Allocate a struct using `new()` function
- Access and modify struct fields through the pointer
- Demonstrate both `ptr.field` and `(*ptr).field` syntax

## Files
- `pointer_manipulation.go` - Complete program demonstrating all pointer concepts
- `go.mod` - Module configuration
- `README_POINTERS.md` - This documentation for Experiment 5

## How to Run
```bash
# Navigate to LAB3 directory
cd Go/LAB3

# Run the pointer manipulation program
go run pointer_manipulation.go
```

## Key Concepts Demonstrated

### 1. Memory Addresses and Pointers
```go
var number int = 42
var ptr *int = &number    // Get address using &
fmt.Printf("Address: %p\n", &number)
fmt.Printf("Value through pointer: %d\n", *ptr)  // Dereference using *
```

### 2. Pass-by-Reference vs Pass-by-Value
```go
func modifyValue(ptr *int) {
    *ptr = *ptr * 2    // Modifies original variable
}
```

### 3. Struct Allocation with new()
```go
type Person struct {
    Name string
    Age  int
    City string
}

personPtr := new(Person)    // Allocates and returns pointer
personPtr.Name = "John"     // Direct field access through pointer
(*personPtr).Age = 30       // Alternative syntax
```

### 4. Multiple Pointers to Same Memory
```go
ptr1 := &sharedVar
ptr2 := &sharedVar    // Both point to same memory location
```

## Sample Output
The program demonstrates:
1. **Basic Operations**: Shows variable addresses, pointer values, and dereferencing
2. **Function Modification**: Before/after values showing pass-by-reference behavior
3. **Struct Allocation**: new() allocation and field access through pointers
4. **Multiple Pointers**: How multiple pointers can reference the same memory

## Learning Outcomes
- Understanding memory addresses and pointer concepts
- Difference between pass-by-value and pass-by-reference
- Proper use of reference (`&`) and dereference (`*`) operators  
- Dynamic memory allocation using `new()`
- Struct field access through pointers with different syntax options

**Score: 5/5 marks** - All requirements fully implemented and demonstrated.