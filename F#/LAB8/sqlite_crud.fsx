#r "nuget: Microsoft.Data.Sqlite, 8.0.34"

open Microsoft.Data.Sqlite

type Student =
    { Id: int
      Name: string
      Age: int }

let connection = new SqliteConnection("Data Source=:memory:")
connection.Open()

let executeNonQuery (sql: string) (configure: SqliteCommand -> unit) =
    use command = connection.CreateCommand()
    command.CommandText <- sql
    configure command
    command.ExecuteNonQuery() |> ignore

let printStudents (operation: string) =
    use command = connection.CreateCommand()
    command.CommandText <- "SELECT Id, Name, Age FROM Students ORDER BY Id;"

    use reader = command.ExecuteReader()
    let students = ResizeArray<Student>()

    while reader.Read() do
        students.Add(
            { Id = reader.GetInt32(0)
              Name = reader.GetString(1)
              Age = reader.GetInt32(2) }
        )

    printfn "%s" operation

    if students.Count = 0 then
        printfn "  Students: (none)"
    else
        students
        |> Seq.iter (fun student -> printfn "  Id=%d, Name=%s, Age=%d" student.Id student.Name student.Age)

executeNonQuery
    "CREATE TABLE Students (Id INTEGER PRIMARY KEY, Name TEXT NOT NULL, Age INTEGER NOT NULL);"
    ignore

executeNonQuery
    "INSERT INTO Students (Id, Name, Age) VALUES ($id, $name, $age);"
    (fun command ->
        command.Parameters.AddWithValue("$id", 1) |> ignore
        command.Parameters.AddWithValue("$name", "Ada Lovelace") |> ignore
        command.Parameters.AddWithValue("$age", 20) |> ignore)

printStudents "After Create:"

printStudents "After Read:"

executeNonQuery
    "UPDATE Students SET Name = $name, Age = $age WHERE Id = $id;"
    (fun command ->
        command.Parameters.AddWithValue("$name", "Ada Byron") |> ignore
        command.Parameters.AddWithValue("$age", 21) |> ignore
        command.Parameters.AddWithValue("$id", 1) |> ignore)

printStudents "After Update:"

executeNonQuery
    "DELETE FROM Students WHERE Id = $id;"
    (fun command -> command.Parameters.AddWithValue("$id", 1) |> ignore)

printStudents "After Delete:"

connection.Dispose()
