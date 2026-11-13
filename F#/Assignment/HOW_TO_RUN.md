# How to Run the Interactive Library Management System

## 🚀 Quick Start

### Step 1: Open Terminal/Command Prompt
Navigate to the project directory:
```bash
cd G:\NAPL\F#\Assignment\LibrarySystem
```

### Step 2: Run the Program
```bash
dotnet run
```

### Step 3: You'll see this menu:
```
╔════════════════════════════════════════════════════════╗
║           Library Management System - Menu             ║
╠════════════════════════════════════════════════════════╣
║  1. Add a new book                                     ║
║  2. View all books                                     ║
║  3. Search book by ID                                  ║
║  4. Update a book                                      ║
║  5. Delete a book                                      ║
║  6. Checkout a book                                    ║
║  7. Return a book                                      ║
║  8. View available books                               ║
║  9. View library statistics                            ║
║  0. Exit                                               ║
╚════════════════════════════════════════════════════════╝
```

## 📖 Usage Examples

### Example 1: Add a Book
```
Enter your choice (0-9): 1

=== Add New Book ===
Enter Book ID: 101
Enter Title: The Great Gatsby
Enter Author: F. Scott Fitzgerald
Enter ISBN: 978-0743273565
Enter Publication Year: 1925
✅ Book added successfully!

Press Enter to continue...
```

### Example 2: View All Books
```
Enter your choice (0-9): 2

=== All Books ===
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
ID:     101
Title:  The Great Gatsby
Author: F. Scott Fitzgerald
ISBN:   978-0743273565
Year:   1925
Status: Available
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Total books: 1

Press Enter to continue...
```

### Example 3: Checkout a Book
```
Enter your choice (0-9): 6

=== Checkout Book ===
Enter Book ID to checkout: 101
✅ Book checked out successfully!

Press Enter to continue...
```

### Example 4: View Statistics
```
Enter your choice (0-9): 9

=== Library Statistics ===
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Total Books:       5
Available:         3
Checked Out:       2
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

📚 Books by Author:
  F. Scott Fitzgerald: 1 book(s)
  George Orwell: 2 book(s)

Press Enter to continue...
```

### Example 5: Update a Book
```
Enter your choice (0-9): 4

=== Update Book ===
Enter Book ID to update: 101

Current details:
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
ID:     101
Title:  The Great Gatsby
Author: F. Scott Fitzgerald
ISBN:   978-0743273565
Year:   1925
Status: Available
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Enter new details (press Enter to keep existing):
Title [The Great Gatsby]: 
Author [F. Scott Fitzgerald]: 
ISBN [978-0743273565]: 978-1234567890
Year [1925]: 

✅ Book updated successfully!

Press Enter to continue...
```

## 🎯 Common Tasks

### Adding Multiple Books Quickly
```
Choice: 1 → Add Book 1 → Enter
Choice: 1 → Add Book 2 → Enter
Choice: 1 → Add Book 3 → Enter
Choice: 2 → View all books
```

### Managing Book Status
```
Choice: 6 → Checkout book (makes it unavailable)
Choice: 8 → View available books (see what's in stock)
Choice: 7 → Return book (makes it available again)
```

### Deleting Books
```
Choice: 3 → Search for book (verify it exists)
Choice: 5 → Delete book → Confirm with 'y'
```

## ⚙️ Features

✅ **Custom Input** - All data entered by you  
✅ **Input Validation** - Won't accept invalid data  
✅ **Persistent Storage** - Data saved in SQLite database  
✅ **Easy Updates** - Press Enter to keep current values  
✅ **Confirmation** - Asks before deleting  
✅ **Statistics** - Real-time library analytics  

## 💾 Database

- **Location**: `library.db` (created automatically)
- **Persistent**: Data saved between runs
- **Reset**: Delete `library.db` to start fresh

## 🔄 To Exit

Type `0` and press Enter at any time to exit the program.

## 🎨 Interactive vs Demo Version

**Two versions available:**

1. **Interactive** (`ProgramInteractive.fs`) - You enter all data ✅ Currently active
2. **Demo** (`Program.fs`) - Pre-loaded with sample data

To switch back to demo version, edit `LibrarySystem.fsproj` and change:
```xml
<Compile Include="ProgramInteractive.fs" />
```
to:
```xml
<Compile Include="Program.fs" />
```

Happy managing your library! 📚✨
