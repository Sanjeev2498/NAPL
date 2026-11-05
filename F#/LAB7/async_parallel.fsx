open System

let runTask (taskName: string) (delayMilliseconds: int) (result: string) : Async<string> =
    async {
        printfn "%s started." taskName
        do! Async.Sleep delayMilliseconds
        printfn "%s finished after %d ms." taskName delayMilliseconds
        return result
    }

let tasks =
    [| runTask "Task 1" 900 "Result 1"
       runTask "Task 2" 300 "Result 2"
       runTask "Task 3" 600 "Result 3" |]

printfn "Starting all tasks in parallel..."
let results = tasks |> Async.Parallel |> Async.RunSynchronously

printfn "\nAll tasks completed. Results:"
results |> Array.iter (printfn "%s")
