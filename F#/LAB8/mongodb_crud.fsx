#r "nuget: MongoDB.Driver, 3.13.0"

open System
open MongoDB.Bson
open MongoDB.Driver

[<AllowNullLiteral>]
type Student() =
    member val Id = ObjectId.Empty with get, set
    member val Name = "" with get, set
    member val Age = 0 with get, set

let configuredUri = Environment.GetEnvironmentVariable("MONGODB_URI")
let connectionString =
    if String.IsNullOrWhiteSpace(configuredUri) then
        "mongodb://localhost:27017"
    else
        configuredUri

let client = new MongoClient(connectionString)
let database = client.GetDatabase("lab8")
let students = database.GetCollection<Student>("students")
let studentId = ObjectId.GenerateNewId()
let studentFilter = Builders<Student>.Filter.Eq<ObjectId>("_id", studentId)

let printStudent (operation: string) =
    let student = students.Find(studentFilter).FirstOrDefault()
    printfn "%s" operation

    if isNull student then
        printfn "  Student: (none)"
    else
        printfn "  Id=%O, Name=%s, Age=%d" student.Id student.Name student.Age

let newStudent = Student(Id = studentId, Name = "Ada Lovelace", Age = 20)
students.InsertOne(newStudent)
printStudent "After Create:"

printStudent "After Read:"

let update = Builders<Student>.Update.Set("Name", "Ada Byron").Set("Age", 21)
students.UpdateOne(studentFilter, update) |> ignore
printStudent "After Update:"

students.DeleteOne(studentFilter) |> ignore
printStudent "After Delete:"
