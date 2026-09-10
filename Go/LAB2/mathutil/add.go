package mathutil

import "strings"

func Add(num1, num2 int) int {
	sum := num1 + num2
	return sum
}

func Reverse(text string) string {
	runes := []rune(text)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func CountVowels(text string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, ch := range text {
		if strings.ContainsRune(vowels, ch) {
			count++
		}
	}
	return count
}

func Factorial(n int) int {
	if n < 0 {
		return 0
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func Power(base, exponent int) int {
	if exponent < 0 {
		return 0
	}
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}