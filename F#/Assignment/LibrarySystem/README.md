# F# Library Management System - Assignment

## Overview
A simple library management system demonstrating F# domain modeling, data persistence, and functional programming concepts.

## Requirements Fulfilled

### 1. Domain Model ✅
- **Book record type** with fields: Id, Title, Author, ISBN, Year, Available
- Clean F# types and records
- Helper functions using pattern matching

### 2. Higher-Order Functions ✅
- **filter**: Get available books, books by author, books by year
- **map**: Convert books to summary strings
- **Pattern matching**: Book status, availability checks

### 3. Data Persistence ✅
- SQLite database for storage
- **CRUD Operations**:
  - **CREATE**: Add books to database
  - **READ**: Get all books, get book by ID
  - **UPDATE**: Update book details, checkout/return books
  - **DELETE**: Remove books from database
- **Async workflows** for all database operations

## Files
- `Library.fs` - Domain model and business logic
- `Database.fs` - Database operations with async workflows
- `Program.fs` - Main program demonstrating all features
- `LibrarySystem.fsproj` - Project configuration
- `README.md` - Documentation

## How to Run

### Build the project
```bash
cd "F#/Assignment/LibrarySystem"
dotnet build
```

### Run the program
```bash
dotnet run
```

## Key Features

### Domain Model
```fsharp
type Book = {
    Id: int
    Title: string
    Author: string
    ISBN: string
    Year: int
    Available: bool
}
```

### Higher-Order Functions
```fsharp
// Filter available books
let getAvailableBooks books =
    books |> List.filter isAvailable

// Map to summaries
let getBookSummaries books =
    books |> List.map (fun book -> sprintf "...")
```

### Pattern Matching
```fsharp
let getBookStatus book =
    match book.Available with
    | true -> "Available"
    | false -> "Checked Out"
```

### Async Database Operations
```fsharp
let addBook book =
    async {
        use connection = new SQLiteConnection(connectionString)
        // ... database operations
        return result
    }
```

## F# Concepts Demonstrated
✅ Records and types  
✅ Pattern matching  
✅ Higher-order functions (map, filter)  
✅ Async workflows  
✅ Database operations (CRUD)  
✅ Functional composition  
✅ Type safety  
✅ Immutability  

**Assignment completed successfully!**
