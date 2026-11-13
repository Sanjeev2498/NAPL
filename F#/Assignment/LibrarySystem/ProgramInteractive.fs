open System
open Library
open Database

// Helper function to read input
let readInput prompt =
    printf "%s" prompt
    Console.ReadLine()

let readInt prompt =
    let rec loop () =
        printf "%s" prompt
        match Int32.TryParse(Console.ReadLine()) with
        | true, value -> value
        | false, _ -> 
            printfn "Invalid input. Please enter a number."
            loop()
    loop()

let readBool prompt =
    let rec loop () =
        printf "%s (y/n): " prompt
        match Console.ReadLine().ToLower() with
        | "y" | "yes" -> true
        | "n" | "no" -> false
        | _ -> 
            printfn "Invalid input. Please enter y or n."
            loop()
    loop()

// Display menu
let displayMenu () =
    printfn ""
    printfn "╔════════════════════════════════════════════════════════╗"
    printfn "║           Library Management System - Menu             ║"
    printfn "╠════════════════════════════════════════════════════════╣"
    printfn "║  1. Add a new book                                     ║"
    printfn "║  2. View all books                                     ║"
    printfn "║  3. Search book by ID                                  ║"
    printfn "║  4. Update a book                                      ║"
    printfn "║  5. Delete a book                                      ║"
    printfn "║  6. Checkout a book                                    ║"
    printfn "║  7. Return a book                                      ║"
    printfn "║  8. View available books                               ║"
    printfn "║  9. View library statistics                            ║"
    printfn "║  0. Exit                                               ║"
    printfn "╚════════════════════════════════════════════════════════╝"
    printfn ""

// Add book
let addBookInteractive () =
    printfn "\n=== Add New Book ==="
    let id = readInt "Enter Book ID: "
    let title = readInput "Enter Title: "
    let author = readInput "Enter Author: "
    let isbn = readInput "Enter ISBN: "
    let year = readInt "Enter Publication Year: "
    
    let book = createBook id title author isbn year
    let result = addBook book |> Async.RunSynchronously
    
    if result then
        printfn "✅ Book added successfully!"
    else
        printfn "❌ Failed to add book. ID might already exist."

// View all books
let viewAllBooks () =
    printfn "\n=== All Books ==="
    let books = getAllBooks() |> Async.RunSynchronously
    
    if List.isEmpty books then
        printfn "No books in the library."
    else
        displayBooks books
        printfn "Total books: %d" (List.length books)

// Search by ID
let searchById () =
    printfn "\n=== Search Book by ID ==="
    let id = readInt "Enter Book ID: "
    
    let bookOption = getBookById id |> Async.RunSynchronously
    match bookOption with
    | Some book -> displayBook book
    | None -> printfn "❌ Book not found with ID: %d" id

// Update book
let updateBookInteractive () =
    printfn "\n=== Update Book ==="
    let id = readInt "Enter Book ID to update: "
    
    let bookOption = getBookById id |> Async.RunSynchronously
    match bookOption with
    | None -> printfn "❌ Book not found with ID: %d" id
    | Some existingBook ->
        printfn "\nCurrent details:"
        displayBook existingBook
        
        printfn "\nEnter new details (press Enter to keep existing):"
        
        let title = 
            let input = readInput (sprintf "Title [%s]: " existingBook.Title)
            if String.IsNullOrWhiteSpace(input) then existingBook.Title else input
        
        let author = 
            let input = readInput (sprintf "Author [%s]: " existingBook.Author)
            if String.IsNullOrWhiteSpace(input) then existingBook.Author else input
        
        let isbn = 
            let input = readInput (sprintf "ISBN [%s]: " existingBook.ISBN)
            if String.IsNullOrWhiteSpace(input) then existingBook.ISBN else input
        
        let yearInput = readInput (sprintf "Year [%d]: " existingBook.Year)
        let year = 
            if String.IsNullOrWhiteSpace(yearInput) then 
                existingBook.Year 
            else 
                match Int32.TryParse(yearInput) with
                | true, y -> y
                | false, _ -> existingBook.Year
        
        let updatedBook = { existingBook with Title = title; Author = author; ISBN = isbn; Year = year }
        let result = updateBook updatedBook |> Async.RunSynchronously
        
        if result then
            printfn "✅ Book updated successfully!"
        else
            printfn "❌ Failed to update book."

// Delete book
let deleteBookInteractive () =
    printfn "\n=== Delete Book ==="
    let id = readInt "Enter Book ID to delete: "
    
    let bookOption = getBookById id |> Async.RunSynchronously
    match bookOption with
    | None -> printfn "❌ Book not found with ID: %d" id
    | Some book ->
        displayBook book
        let confirm = readBool "Are you sure you want to delete this book?"
        
        if confirm then
            let result = deleteBook id |> Async.RunSynchronously
            if result then
                printfn "✅ Book deleted successfully!"
            else
                printfn "❌ Failed to delete book."
        else
            printfn "Delete cancelled."

// Checkout book
let checkoutBookInteractive () =
    printfn "\n=== Checkout Book ==="
    let id = readInt "Enter Book ID to checkout: "
    
    let bookOption = getBookById id |> Async.RunSynchronously
    match bookOption with
    | None -> printfn "❌ Book not found with ID: %d" id
    | Some book ->
        if not book.Available then
            printfn "❌ Book is already checked out."
        else
            let checkedOutBook = checkoutBook book
            let result = updateBook checkedOutBook |> Async.RunSynchronously
            
            if result then
                printfn "✅ Book checked out successfully!"
            else
                printfn "❌ Failed to checkout book."

// Return book
let returnBookInteractive () =
    printfn "\n=== Return Book ==="
    let id = readInt "Enter Book ID to return: "
    
    let bookOption = getBookById id |> Async.RunSynchronously
    match bookOption with
    | None -> printfn "❌ Book not found with ID: %d" id
    | Some book ->
        if book.Available then
            printfn "❌ Book is already available (not checked out)."
        else
            let returnedBook = returnBook book
            let result = updateBook returnedBook |> Async.RunSynchronously
            
            if result then
                printfn "✅ Book returned successfully!"
            else
                printfn "❌ Failed to return book."

// View available books
let viewAvailableBooks () =
    printfn "\n=== Available Books ==="
    let allBooks = getAllBooks() |> Async.RunSynchronously
    let availableBooks = getAvailableBooks allBooks
    
    if List.isEmpty availableBooks then
        printfn "No available books."
    else
        displayBooks availableBooks
        printfn "Available books: %d" (List.length availableBooks)

// View statistics
let viewStatistics () =
    printfn "\n=== Library Statistics ==="
    let books = getAllBooks() |> Async.RunSynchronously
    let totalBooks = List.length books
    let availableBooks = getAvailableBooks books
    let availableCount = List.length availableBooks
    let checkedOutCount = totalBooks - availableCount
    
    printfn "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    printfn "Total Books:       %d" totalBooks
    printfn "Available:         %d" availableCount
    printfn "Checked Out:       %d" checkedOutCount
    printfn "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    
    if totalBooks > 0 then
        printfn "\n📚 Books by Author:"
        books 
        |> List.groupBy (fun b -> b.Author)
        |> List.iter (fun (author, bookList) -> 
            printfn "  %s: %d book(s)" author (List.length bookList))

// Main loop
let rec mainLoop () =
    displayMenu()
    let choice = readInput "Enter your choice (0-9): "
    
    match choice with
    | "1" -> addBookInteractive()
    | "2" -> viewAllBooks()
    | "3" -> searchById()
    | "4" -> updateBookInteractive()
    | "5" -> deleteBookInteractive()
    | "6" -> checkoutBookInteractive()
    | "7" -> returnBookInteractive()
    | "8" -> viewAvailableBooks()
    | "9" -> viewStatistics()
    | "0" -> 
        printfn "\nThank you for using Library Management System!"
        printfn "Goodbye! 👋"
        ()
    | _ -> 
        printfn "❌ Invalid choice. Please enter a number between 0 and 9."
        mainLoop()
    
    if choice <> "0" then
        printfn "\nPress Enter to continue..."
        Console.ReadLine() |> ignore
        mainLoop()

[<EntryPoint>]
let main argv =
    printfn "╔════════════════════════════════════════════════════════╗"
    printfn "║   Library Management System (F# - Interactive)         ║"
    printfn "╚════════════════════════════════════════════════════════╝"
    printfn ""
    
    // Initialize database
    printfn "Initializing database..."
    initializeDatabase() |> Async.RunSynchronously
    printfn "✅ Database ready!"
    
    // Start interactive loop
    mainLoop()
    
    0 // Exit code
