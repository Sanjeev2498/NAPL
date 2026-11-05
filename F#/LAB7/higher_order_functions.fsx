let testLists =
    [ "List 1", [ 1; 2; 3; 4; 5 ]
      "List 2", [ 6; 7; 8; 9; 10 ] ]

printfn "List.map - square each number:"
testLists
|> List.iter (fun (name, numbers) ->
    let squares = List.map (fun number -> number * number) numbers
    printfn "%s: %A -> %A" name numbers squares)

printfn "\nList.filter - keep even numbers:"
testLists
|> List.iter (fun (name, numbers) ->
    let evens = List.filter (fun number -> number % 2 = 0) numbers
    printfn "%s: %A -> %A" name numbers evens)

printfn "\nList.fold - calculate the sum:"
testLists
|> List.iter (fun (name, numbers) ->
    let sum = List.fold (fun total number -> total + number) 0 numbers
    printfn "%s: %A -> %d" name numbers sum)
