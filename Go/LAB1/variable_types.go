package main

import (
	"fmt"
	"unicode/utf8"
)

func variableTypes() {
	intNum := 3232
	intNum = intNum + 1
	fmt.Println(intNum)

	var mystring string = "Hello" + "Sam"
	fmt.Println(mystring)

	fmt.Println(utf8.RuneCountInString("y"))
	var myRune rune = 'a'
	fmt.Println(myRune)
}
