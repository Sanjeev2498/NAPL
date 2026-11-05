# LAB 7: F# Higher-Order Functions and Async Workflows

This lab contains two standalone F# scripts:

1. `async_parallel.fsx` runs three delayed asynchronous tasks concurrently, prints their start and completion messages, and displays all results after they finish.
2. `higher_order_functions.fsx` demonstrates `List.map`, `List.filter`, and `List.fold` with lambda expressions on two lists each.

## Run

Install the .NET SDK, then run each script from this directory:

```powershell
dotnet fsi .\async_parallel.fsx
dotnet fsi .\higher_order_functions.fsx
```

In the asynchronous workflow output, tasks finish in a different order than they started because their simulated delays differ.
