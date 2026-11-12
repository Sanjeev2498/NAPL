# LAB 8: OOP and Database CRUD in F#

This lab contains three standalone F# scripts:

1. `oop_classes.fsx` defines an abstract `Shape` class and overrides its virtual `Area` method in `Circle` and `Rectangle`, then calls the method polymorphically.
2. `sqlite_crud.fsx` performs Create, Read, Update, and Delete operations on a `Student` table in an in-memory SQLite database. It prints the table contents after each operation.
3. `mongodb_crud.fsx` performs the same operations on a MongoDB `students` collection. It prints the state of the student used by the script after each operation.

## Run

Install the .NET SDK, then run each script from this directory:

```powershell
dotnet fsi .\oop_classes.fsx
dotnet fsi .\sqlite_crud.fsx
dotnet fsi .\mongodb_crud.fsx
```

The database scripts restore their NuGet dependencies when run. SQLite runs entirely in memory. MongoDB requires a reachable MongoDB server at `mongodb://localhost:27017` by default; set the `MONGODB_URI` environment variable to use another connection string.
