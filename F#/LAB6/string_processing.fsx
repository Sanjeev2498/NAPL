open System
open System.Text.RegularExpressions

let reverseString (value: string) =
    value |> Seq.rev |> Seq.toArray |> fun characters -> String(characters)

let isPalindrome (value: string) =
    value = reverseString value

let transformString (value: string) =
    if not (String.IsNullOrEmpty value) && (value |> Seq.forall Char.IsLetter) then
        value.ToUpperInvariant()
    elif not (String.IsNullOrEmpty value) && (value |> Seq.forall Char.IsDigit) then
        value + value
    else
        value

let emailPattern = @"^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$"

let isValidEmail (value: string) =
    Regex.IsMatch(value, emailPattern)

printfn "Reverse and palindrome tests:"
[ "level"; "hello"; "abba" ]
|> List.iter (fun value ->
    printfn "\"%s\" -> reverse: \"%s\", palindrome: %b" value (reverseString value) (isPalindrome value))

printfn "\nString transformation tests:"
[ "hello"; "1234"; "hi123" ]
|> List.iter (fun value ->
    printfn "\"%s\" -> \"%s\"" value (transformString value))

printfn "\nEmail validation tests:"
[ "student@example.com"; "name.surname@college.edu"; "invalid-email" ]
|> List.iter (fun value ->
    printfn "\"%s\" -> %s" value (if isValidEmail value then "valid" else "invalid"))
