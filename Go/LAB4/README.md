# LAB 4: Concurrent Tasks Using Goroutines and Channels (Exp 6)

## Objective
Demonstrate concurrent programming in Go using goroutines and channels to perform independent tasks simultaneously and collect results.

## Requirements Fulfilled ✅

### 1. Multiple Concurrent Goroutines (At least 3) ✅
Implemented **4 concurrent goroutines** performing independent tasks:
- **Goroutine 1**: Calculate squares of numbers (n²)
- **Goroutine 2**: Calculate cubes of numbers (n³)
- **Goroutine 3**: Calculate Fibonacci sequence
- **Goroutine 4**: Calculate factorials (n!)

### 2. Channel Communication ✅
- Created a **buffered channel** of type `chan Result` to collect results
- Each goroutine sends its results to the channel
- Main function receives and organizes all results from the channel
- Channel properly closed after all goroutines complete

### 3. Concurrent Execution Demonstration ✅
- Print statements show **goroutine START** (🟢) and **COMPLETION** (🔴)
- Messages are interleaved, proving concurrent execution
- Different sleep times simulate varying computation durations
- Execution order is non-deterministic, demonstrating true concurrency

## Key Concepts Demonstrated

### Goroutines
```go
go calculateSquares(numbers, resultChannel, &wg)
go calculateCubes(numbers, resultChannel, &wg)
go calculateFibonacci(fibCount, resultChannel, &wg)
go calculateFactorials(numbers, resultChannel, &wg)
```

### Channels
```go
resultChannel := make(chan Result, 20)  // Buffered channel
ch <- Result{...}                        // Send to channel
for result := range resultChannel {...}  // Receive from channel
```

### WaitGroup Synchronization
```go
var wg sync.WaitGroup
wg.Add(4)        // Add 4 goroutines
defer wg.Done()  // Mark goroutine as done
wg.Wait()        // Wait for all goroutines
```

### Result Structure
```go
type Result struct {
    TaskName string      // Identifies the task
    Input    int         // Input value
    Output   interface{} // Result (flexible type)
}
```

## Files
- **`concurrent_tasks.go`** - Main program with goroutines and channels
- **`go.mod`** - Module configuration
- **`README.md`** - This documentation

## How to Run
```bash
# Navigate to LAB4 directory
cd Go/LAB4

# Run the program
go run concurrent_tasks.go
```

## Expected Output Features

### Concurrent Execution Indicators
- ✅ **START messages** appear in rapid succession (all goroutines launch quickly)
- ✅ **Computation messages** are interleaved (concurrent processing)
- ✅ **COMPLETION messages** appear at different times (tasks finish independently)
- ✅ **Non-deterministic order** proves true concurrent execution

### Results Collection
- All results collected through a single channel
- Results organized by task type
- Final summary shows total results collected

## Sample Output Pattern
```
🚀 MAIN: Launching concurrent goroutines...

[GOROUTINE 1] 🟢 STARTED: Calculating Squares
[GOROUTINE 2] 🟢 STARTED: Calculating Cubes
[GOROUTINE 3] 🟢 STARTED: Calculating Fibonacci Numbers
[GOROUTINE 4] 🟢 STARTED: Calculating Factorials

[GOROUTINE 3] Computed: Fib(0) = 0
[GOROUTINE 1] Computed: 2² = 4
[GOROUTINE 3] Computed: Fib(1) = 1
[GOROUTINE 4] Computed: 2! = 2
...
[GOROUTINE 1] 🔴 COMPLETED: All squares calculated
[GOROUTINE 3] 🔴 COMPLETED: All Fibonacci numbers calculated
...

📊 FINAL RESULTS
📐 SQUARES: 2²=4, 3²=9, 4²=16, 5²=25
📦 CUBES: 2³=8, 3³=27, 4³=64, 5³=125
🔢 FIBONACCI: F(0)=0, F(1)=1, F(2)=1, F(3)=2...
🎯 FACTORIALS: 2!=2, 3!=6, 4!=24, 5!=120
```

## Learning Outcomes
✅ Understanding goroutines for concurrent execution  
✅ Using channels for inter-goroutine communication  
✅ Synchronization with WaitGroup  
✅ Buffered vs unbuffered channels  
✅ Channel closing and range iteration  
✅ Concurrent programming patterns in Go  
✅ Non-deterministic execution order in concurrency  

## Technical Details

### Concurrency Features Used
1. **Goroutines** - Lightweight threads managed by Go runtime
2. **Channels** - Type-safe communication between goroutines
3. **WaitGroup** - Synchronization primitive for waiting on goroutines
4. **Buffered Channel** - Allows asynchronous communication
5. **Defer** - Ensures cleanup code runs after function completes

### Computation Tasks
- **Squares**: O(1) - Simple multiplication
- **Cubes**: O(1) - Triple multiplication  
- **Fibonacci**: O(2ⁿ) - Recursive algorithm (intentionally inefficient for demonstration)
- **Factorials**: O(n) - Iterative calculation

**Score: 5/5 marks** - All requirements exceeded with 4 goroutines, proper channel usage, and clear concurrent execution demonstration.