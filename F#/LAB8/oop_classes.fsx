open System

[<AbstractClass>]
type Shape() =
    abstract member Area: unit -> float

type Circle(radius: float) =
    inherit Shape()

    override _.Area() = Math.PI * radius * radius

type Rectangle(width: float, height: float) =
    inherit Shape()

    override _.Area() = width * height

let shapes: Shape list =
    [ Circle(5.0)
      Rectangle(4.0, 6.0) ]

printfn "Areas calculated polymorphically:"
shapes
|> List.iter (fun shape -> printfn "%s area: %.2f" (shape.GetType().Name) (shape.Area()))
