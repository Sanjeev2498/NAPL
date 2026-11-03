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

	runCollectionManagement()
}

func runCollectionManagement() {
	fmt.Println("\nCollection Management Using Slices and Maps - Exp 3")

	students := []string{"Alice", "Bob"}
	fmt.Println("\nInitial students:", students)

	students = append(students, "Charlie")
	fmt.Println("After adding Charlie:", students)

	students = append(students[:1], students[2:]...)
	fmt.Println("After removing student at index 1:", students)

	students[1] = "Diana"
	fmt.Println("After updating student at index 1:", students)

	marks := make(map[string]int)
	fmt.Println("\nInitial subject marks:", marks)

	marks["Mathematics"] = 95
	fmt.Println("After inserting Mathematics:", marks)

	marks["Science"] = 88
	fmt.Println("After inserting Science:", marks)

	delete(marks, "Mathematics")
	fmt.Println("After deleting Mathematics:", marks)

	mark := marks["Science"]
	fmt.Println("Lookup Science:", mark)
	fmt.Println("Map after lookup:", marks)
}
