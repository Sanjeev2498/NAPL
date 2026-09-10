package main

import (
	"fmt"
	"strconv"
)

func main() {
    num := 123
    s := strconv.Itoa(num)

    fmt.Println(s)      // 123
    fmt.Printf("%T\n", s) // string
}