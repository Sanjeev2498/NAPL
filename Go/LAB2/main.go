package main

import (
	"fmt"

	"MyProject/mathutil"
)

func main() {
	text := "GoLang"
	fmt.Println("Reverse:", mathutil.Reverse(text))
	fmt.Println("Vowels:", mathutil.CountVowels("Hello, World!"))
	fmt.Println("Factorial:", mathutil.Factorial(5))
	fmt.Println("Power:", mathutil.Power(2, 5))
	fmt.Println("Add:", mathutil.Add(10, 7))
}
