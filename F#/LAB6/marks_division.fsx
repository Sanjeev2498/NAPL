open System

let readMark subject =
    printf "Enter marks for %s (0-100): " subject

    match Int32.TryParse(Console.ReadLine()) with
    | true, mark when mark >= 0 && mark <= 100 -> float mark
    | true, _ -> failwith "Marks must be between 0 and 100."
    | false, _ -> failwith "Please enter a valid whole number."

let marks =
    [ "Subject 1"; "Subject 2"; "Subject 3"; "Subject 4" ]
    |> List.map readMark

let percentage = List.average marks

let division =
    match percentage with
    | value when value >= 60.0 -> "First Division"
    | value when value >= 50.0 -> "Second Division"
    | value when value >= 40.0 -> "Third Division"
    | _ -> "Fail"

printfn "Percentage: %.2f%%" percentage
printfn "Classification: %s" division
