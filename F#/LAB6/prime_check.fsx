open System

let isPrime number =
    if number < 2 then
        false
    else
        let rec hasDivisor divisor =
            if divisor > number / divisor then
                false
            elif number % divisor = 0 then
                true
            else
                hasDivisor (divisor + 1)

        not (hasDivisor 2)

let testInputs = [ 2; 9; 17; 1; -5 ]

printfn "Prime number test cases:"

testInputs
|> List.iter (fun number ->
    printfn "%d: %s" number (if isPrime number then "Prime" else "Not prime"))

printf "Enter an integer to check: "

match Int32.TryParse(Console.ReadLine()) with
| true, number ->
    printfn "%d is %s." number (if isPrime number then "prime" else "not prime")
| false, _ ->
    eprintfn "Please enter a valid integer."
    Environment.ExitCode <- 1
