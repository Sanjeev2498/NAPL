open System
open Library
open Database

[<EntryPoint>]
let main argv =
    printfn "╔════════════════════════════════════════════════════════╗"
    printfn "║        Library Management System (F#)                  ║"
    printfn "╚════════════════════════════════════════════════════════╝"
    printfn ""

    // Initialize database
    printfn "Initializing database..."
    initializeDatabase() |> Async.RunSynchronously
    
    // Clear existing data for clean demo
    clearAllBooks() |> Async.RunSynchronously
    printfn "Database ready!"
    printfn ""

    // Create sample books
    let books = [
        createBook 1 "The Great Gatsby" "F. Scott Fitzgerald" "978-0743273565" 1925
        createBook 2 "1984" "George Orwell" "978-0451524935" 1949
        createBook 3 "To Kill a Mockingbird" "Harper Lee" "978-0061120084" 1960
        createBook 4 "Pride and Prejudice" "Jane Austen" "978-0141439518" 1813
        createBook 5 "The Hobbit" "J.R.R. Tolkien" "978-0547928227" 1937
    ]

    // CRUD Operations Demo
    printfn "═══════════════════════════════════════════════════════"
    printfn "1. CREATE - Adding books to database..."
    printfn "═══════════════════════════════════════════════════════"
    
    books 
    |> List.iter (fun book -> 
        let result = addBook book |> Async.RunSynchronously
        printfn "Added: %s by %s" book.Title book.Author)
    
    printfn ""

    // READ - Get all books
    printfn "═══════════════════════════════════════════════════════"
    printfn "2. READ - Retrieving all books from database..."
    printfn "═══════════════════════════════════════════════════════"
    
    let allBooks = getAllBooks() |> Async.RunSynchronously
    displayBooks allBooks
    printfn ""

    // Pattern Matching & Higher-Order Functions
    printfn "═══════════════════════════════════════════════════════"
    printfn "3. FILTER - Using higher-order functions..."
    printfn "═══════════════════════════════════════════════════════"
    
    printfn "\n📚 Books by George Orwell:"
    let orwellBooks = getBooksByAuthor "George Orwell" allBooks
    orwellBooks |> getBookSummaries |> List.iter (printfn "  %s")
    
    printfn "\n📅 Books published after 1940:"
    let modernBooks = getBooksAfterYear 1940 allBooks
    modernBooks |> getBookSummaries |> List.iter (printfn "  %s")
    
    printfn ""

    // UPDATE - Checkout a book
    printfn "═══════════════════════════════════════════════════════"
    printfn "4. UPDATE - Checking out a book..."
    printfn "═══════════════════════════════════════════════════════"
    
    let bookToCheckout = allBooks |> List.head
    let checkedOutBook = checkoutBook bookToCheckout
    
    let updateResult = updateBook checkedOutBook |> Async.RunSynchronously
    printfn "Checked out: %s" checkedOutBook.Title
    
    // Verify update
    let updatedBook = getBookById checkedOutBook.Id |> Async.RunSynchronously
    match updatedBook with
    | Some book -> 
        printfn "Status verified: %s" (getBookStatus book)
        displayBook book
    | None -> 
        printfn "Book not found"
    
    printfn ""

    // Show available books using filter
    printfn "═══════════════════════════════════════════════════════"
    printfn "5. FILTER - Available books only..."
    printfn "═══════════════════════════════════════════════════════"
    
    let currentBooks = getAllBooks() |> Async.RunSynchronously
    let availableBooks = getAvailableBooks currentBooks
    
    printfn "\n✅ Available books (%d):" (List.length availableBooks)
    availableBooks |> getBookSummaries |> List.iter (printfn "  %s")
    printfn ""

    // DELETE operation
    printfn "═══════════════════════════════════════════════════════"
    printfn "6. DELETE - Removing a book..."
    printfn "═══════════════════════════════════════════════════════"
    
    let bookIdToDelete = 5
    let deleteResult = deleteBook bookIdToDelete |> Async.RunSynchronously
    
    if deleteResult then
        printfn "Deleted book with ID: %d" bookIdToDelete
    else
        printfn "Book not found"
    
    // Show remaining books
    let remainingBooks = getAllBooks() |> Async.RunSynchronously
    printfn "\nRemaining books: %d" (List.length remainingBooks)
    printfn ""

    // Summary using pattern matching and map
    printfn "═══════════════════════════════════════════════════════"
    printfn "7. SUMMARY - Library Statistics"
    printfn "═══════════════════════════════════════════════════════"
    
    let finalBooks = getAllBooks() |> Async.RunSynchronously
    let totalBooks = List.length finalBooks
    let availableCount = (getAvailableBooks finalBooks) |> List.length
    let checkedOutCount = totalBooks - availableCount
    
    printfn "Total books:      %d" totalBooks
    printfn "Available:        %d" availableCount
    printfn "Checked out:      %d" checkedOutCount
    printfn ""
    
    printfn "All operations completed successfully!"
    printfn "Database file: library.db"
    
    0 // Exit code
