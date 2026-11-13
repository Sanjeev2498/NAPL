module Library

open System

// Domain Model - Book record type
type Book = {
    Id: int
    Title: string
    Author: string
    ISBN: string
    Year: int
    Available: bool
}

// Create a new book
let createBook id title author isbn year =
    {
        Id = id
        Title = title
        Author = author
        ISBN = isbn
        Year = year
        Available = true
    }

// Helper functions using pattern matching
let getBookStatus book =
    match book.Available with
    | true -> "Available"
    | false -> "Checked Out"

let isAvailable book = book.Available

let checkoutBook book =
    { book with Available = false }

let returnBook book =
    { book with Available = true }

// Higher-order functions for processing books

// Filter available books
let getAvailableBooks books =
    books |> List.filter isAvailable

// Filter books by author
let getBooksByAuthor author books =
    books |> List.filter (fun book -> book.Author = author)

// Filter books by year
let getBooksAfterYear year books =
    books |> List.filter (fun book -> book.Year >= year)

// Map books to summary strings
let getBookSummaries books =
    books 
    |> List.map (fun book -> 
        sprintf "%d. %s by %s (%d) - %s" 
            book.Id book.Title book.Author book.Year (getBookStatus book))

// Display a single book
let displayBook book =

    printfn "ID:     %d" book.Id
    printfn "Title:  %s" book.Title
    printfn "Author: %s" book.Author
    printfn "ISBN:   %s" book.ISBN
    printfn "Year:   %d" book.Year
    printfn "Status: %s" (getBookStatus book)
    printfn "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

// Display list of books
let displayBooks books =
    if List.isEmpty books then
        printfn "No books found."
    else
        books |> List.iter displayBook

