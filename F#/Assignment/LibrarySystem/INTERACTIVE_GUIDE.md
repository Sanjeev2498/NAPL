# Interactive Library Management System - User Guide

## How to Run

### Option 1: Using dotnet run
```bash
cd "F#/Assignment/LibrarySystem"
dotnet run
```

### Option 2: Run the executable directly
```bash
cd "F#/Assignment/LibrarySystem/bin/Debug/net10.0"
./LibrarySystem.exe
```

## Features

### Main Menu Options

**1. Add a new book**
   - Enter Book ID (unique number)
   - Enter Title
   - Enter Author
   - Enter ISBN
   - Enter Publication Year

**2. View all books**
   - Displays all books in the library
   - Shows ID, Title, Author, ISBN, Year, Status

**3. Search book by ID**
   - Find a specific book by its ID

**4. Update a book**
   - Modify existing book details
   - Press Enter to keep existing values

**5. Delete a book**
   - Remove a book from the library
   - Requires confirmation

**6. Checkout a book**
   - Mark a book as checked out (unavailable)

**7. Return a book**
   - Mark a checked-out book as available again

**8. View available books**
   - Shows only books that are currently available

**9. View library statistics**
   - Total books count
   - Available vs Checked out
   - Books grouped by author

**0. Exit**
   - Close the application

## Sample Usage Flow

### Adding Your First Book:
```
Enter your choice: 1
Enter Book ID: 101
Enter Title: Harry Potter and the Sorcerer's Stone
Enter Author: J.K. Rowling
Enter ISBN: 978-0439708180
Enter Publication Year: 1997
✅ Book added successfully!
```

### Checking Out a Book:
```
Enter your choice: 6
Enter Book ID to checkout: 101
✅ Book checked out successfully!
```

### Viewing Statistics:
```
Enter your choice: 9

=== Library Statistics ===
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Total Books:       5
Available:         3
Checked Out:       2
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
```

## Tips

- **Book IDs must be unique** - You'll get an error if you try to add a duplicate ID
- **Press Enter to skip** - When updating, press Enter to keep the current value
- **Type carefully** - Invalid inputs will prompt you to re-enter
- **Database persists** - Your books are saved in `library.db` file

## Database Location

The SQLite database file `library.db` is created in the same directory as the executable. Your data persists between runs!

## Common Operations

### Quick Start
1. Run the program
2. Choose option 1 to add some books
3. Choose option 2 to view all books
4. Choose option 6 to checkout a book
5. Choose option 9 to see statistics

### Managing Books
- Add multiple books first
- Use option 3 to search before updating or deleting
- Always view statistics (option 9) to see the current state

Enjoy using the Library Management System! 📚
