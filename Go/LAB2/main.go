package main

import (
	"fmt"
	"sort"
	"strings"

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

	var err error
	students, err = removeAt(students, 1)
	if err != nil {
		fmt.Println("Could not remove student:", err)
		return
	}
	fmt.Println("After removing student at index 1:", students)

	if err := updateAt(students, 1, "Diana"); err != nil {
		fmt.Println("Could not update student:", err)
		return
	}
	fmt.Println("After updating student at index 1:", students)

	marks := make(map[string]int)
	fmt.Println("\nInitial subject marks:", formatMarks(marks))

	marks["Mathematics"] = 95
	fmt.Println("After inserting Mathematics:", formatMarks(marks))

	marks["Science"] = 88
	fmt.Println("After inserting Science:", formatMarks(marks))

	delete(marks, "Mathematics")
	fmt.Println("After deleting Mathematics:", formatMarks(marks))

	mark, found := marks["Science"]
	if found {
		fmt.Printf("Lookup Science: %d\n", mark)
	} else {
		fmt.Println("Lookup Science: not found")
	}
	fmt.Println("Map after lookup:", formatMarks(marks))
}

func removeAt(values []string, index int) ([]string, error) {
	if index < 0 || index >= len(values) {
		return values, fmt.Errorf("index %d is out of range", index)
	}

	return append(values[:index], values[index+1:]...), nil
}

func updateAt(values []string, index int, value string) error {
	if index < 0 || index >= len(values) {
		return fmt.Errorf("index %d is out of range", index)
	}

	values[index] = value
	return nil
}

func formatMarks(marks map[string]int) string {
	subjects := make([]string, 0, len(marks))
	for subject := range marks {
		subjects = append(subjects, subject)
	}
	sort.Strings(subjects)

	entries := make([]string, 0, len(subjects))
	for _, subject := range subjects {
		entries = append(entries, fmt.Sprintf("%s:%d", subject, marks[subject]))
	}
	return "map[" + strings.Join(entries, " ") + "]"
}
