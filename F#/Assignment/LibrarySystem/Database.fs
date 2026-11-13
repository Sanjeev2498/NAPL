module Database

open System.Data.SQLite
open Library

let connectionString = "Data Source=library.db;Version=3;"

// Initialize database with table
let initializeDatabase () =
    async {
        use connection = new SQLiteConnection(connectionString)
        connection.Open()
        
        let createTableQuery = """
            CREATE TABLE IF NOT EXISTS Books (
                Id INTEGER PRIMARY KEY,
                Title TEXT NOT NULL,
                Author TEXT NOT NULL,
                ISBN TEXT NOT NULL,
                Year INTEGER NOT NULL,
                Available INTEGER NOT NULL
            )
        """
        
        use command = new SQLiteCommand(createTableQuery, connection)
        let! _ = command.ExecuteNonQueryAsync() |> Async.AwaitTask
        return ()
    }

// CREATE - Add a book
let addBook book =
    async {
        use connection = new SQLiteConnection(connectionString)
        connection.Open()
        
        let query = """
            INSERT INTO Books (Id, Title, Author, ISBN, Year, Available)
            VALUES (@id, @title, @author, @isbn, @year, @available)
        """
        
        use command = new SQLiteCommand(query, connection)
        command.Parameters.AddWithValue("@id", book.Id) |> ignore
        command.Parameters.AddWithValue("@title", book.Title) |> ignore
        command.Parameters.AddWithValue("@author", book.Author) |> ignore
        command.Parameters.AddWithValue("@isbn", book.ISBN) |> ignore
        command.Parameters.AddWithValue("@year", book.Year) |> ignore
        command.Parameters.AddWithValue("@available", if book.Available then 1 else 0) |> ignore
        
        let! result = command.ExecuteNonQueryAsync() |> Async.AwaitTask
        return result > 0
    }

// READ - Get all books
let getAllBooks () =
    async {
        use connection = new SQLiteConnection(connectionString)
        connection.Open()
        
        let query = "SELECT * FROM Books"
        use command = new SQLiteCommand(query, connection)
        use! reader = command.ExecuteReaderAsync() |> Async.AwaitTask
        
        let books = ResizeArray<Book>()
        
        while reader.Read() do
            let book = {
                Id = reader.GetInt32(0)
                Title = reader.GetString(1)
                Author = reader.GetString(2)
                ISBN = reader.GetString(3)
                Year = reader.GetInt32(4)
                Available = reader.GetInt32(5) = 1
            }
            books.Add(book)
        
        return books |> Seq.toList
    }

// READ - Get book by ID
let getBookById id =
    async {
        use connection = new SQLiteConnection(connectionString)
        connection.Open()
        
        let query = "SELECT * FROM Books WHERE Id = @id"
        use command = new SQLiteCommand(query, connection)
        command.Parameters.AddWithValue("@id", id) |> ignore
        
        use! reader = command.ExecuteReaderAsync() |> Async.AwaitTask
        
        if reader.Read() then
            let book = {
                Id = reader.GetInt32(0)
                Title = reader.GetString(1)
                Author = reader.GetString(2)
                ISBN = reader.GetString(3)
                Year = reader.GetInt32(4)
                Available = reader.GetInt32(5) = 1
            }
            return Some book
        else
            return None
    }

// UPDATE - Update a book
let updateBook book =
    async {
        use connection = new SQLiteConnection(connectionString)
        connection.Open()
        
        let query = """
            UPDATE Books 
            SET Title = @title, Author = @author, ISBN = @isbn, 
                Year = @year, Available = @available
            WHERE Id = @id
        """
        
        use command = new SQLiteCommand(query, connection)
        command.Parameters.AddWithValue("@id", book.Id) |> ignore
        command.Parameters.AddWithValue("@title", book.Title) |> ignore
        command.Parameters.AddWithValue("@author", book.Author) |> ignore
        command.Parameters.AddWithValue("@isbn", book.ISBN) |> ignore
        command.Parameters.AddWithValue("@year", book.Year) |> ignore
        command.Parameters.AddWithValue("@available", if book.Available then 1 else 0) |> ignore
        
        let! result = command.ExecuteNonQueryAsync() |> Async.AwaitTask
        return result > 0
    }

// DELETE - Remove a book
let deleteBook id =
    async {
        use connection = new SQLiteConnection(connectionString)
        connection.Open()
        
        let query = "DELETE FROM Books WHERE Id = @id"
        use command = new SQLiteCommand(query, connection)
        command.Parameters.AddWithValue("@id", id) |> ignore
        
        let! result = command.ExecuteNonQueryAsync() |> Async.AwaitTask
        return result > 0
    }

// Clear all books (for testing)
let clearAllBooks () =
    async {
        use connection = new SQLiteConnection(connectionString)
        connection.Open()
        
        let query = "DELETE FROM Books"
        use command = new SQLiteCommand(query, connection)
        let! _ = command.ExecuteNonQueryAsync() |> Async.AwaitTask
        return ()
    }
